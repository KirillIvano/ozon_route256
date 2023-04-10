package worker_pool

import (
	"context"
	"route256/libs/logger"
	"sync"
)

type WorkerPool struct {
	// входящая работа
	workIn chan func()
	// отслеживаем состояние обработки входящего канала
	wgIn *sync.WaitGroup
	// контекст, в котором был создан пул
	ctx context.Context
	// количество воркеров
	maxConn int32
}

// Закрывает пул, ожидая выполнения всей переданной работы
func (w *WorkerPool) GracefulClose() {
	logger.Info("[worker_pool]: closing, waiting for incoming work to stop ")
	w.wgIn.Wait()

	logger.Info("[worker_pool]: closing, none work left, freeing channel")
	close(w.workIn)
}

// добавляет работу в пул
func (w *WorkerPool) Run(f func()) {
	w.wgIn.Add(1)
	defer w.wgIn.Done()

	select {
	case <-w.ctx.Done():
	case w.workIn <- f:
	}
}

// стартуем корутины-воркеры
func (w *WorkerPool) bootstrap() {
	for i := 0; i < int(w.maxConn); i++ {
		go func() {
			// следим за входным каналом
			for task := range w.workIn {
				// если приходит работа - выполняем ее
				task()
			}
		}()
	}
}

func New(ctx context.Context, maxConn int32) *WorkerPool {
	w := &WorkerPool{
		maxConn: maxConn,
		workIn:  make(chan func(), maxConn),
		wgIn:    &sync.WaitGroup{},
		ctx:     ctx,
	}

	w.bootstrap()

	return w
}
