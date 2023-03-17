package worker_pool

import (
	"context"
	"sync"
)

type WorkerPool struct {
	workIn chan func()

	// отслеживаем состояние обработки входящего канала
	wgIn *sync.WaitGroup
	// отслеживаем состояние выходных каналов
	wgOut *sync.WaitGroup

	// контекст, в котором был создан пул
	ctx context.Context

	maxConn int32
	// текущее количество коннектов
	currConn int32
}

// а че с контекстом делать
func (w *WorkerPool) GracefulClose() {
	w.wgIn.Wait()
	close(w.workIn)

	w.wgOut.Wait()
}

func (w *WorkerPool) Run(f func()) {
	// добавляем в обе очереди ожидание, чтобы закрыть их отдельно
	w.wgIn.Add(1)
	w.wgOut.Add(1)

	go func() {
		select {
		case <-w.ctx.Done():
		case w.workIn <- f:
		}

		w.wgIn.Done()
	}()
}

// стартуем корутины-воркеры
func (w *WorkerPool) bootstrap() {
	for i := 0; i < int(w.maxConn); i++ {
		go func() {
			// следим за входным каналом
			for task := range w.workIn {
				// если приходит работа - выполняем ее
				task()
				// сообщаем, что работа выполнена, чтобы закрыть канал позже
				w.wgOut.Done()
			}
		}()
	}
}

func New(ctx context.Context, maxConn int32) *WorkerPool {
	w := &WorkerPool{
		maxConn: maxConn,
		workIn:  make(chan func(), maxConn),
		wgIn:    &sync.WaitGroup{},
		wgOut:   &sync.WaitGroup{},
		ctx:     ctx,
	}

	w.bootstrap()

	return w
}
