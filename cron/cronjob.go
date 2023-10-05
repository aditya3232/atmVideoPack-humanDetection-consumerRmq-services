package cron

// cron for service DelOneDayOldFiles()

// func init() {
// 	log.Info("Cron Job Started")

// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		cron := cron.New(cron.WithChain(
// 			cron.SkipIfStillRunning(cron.DefaultLogger),
// 		))

// 		cron.AddFunc("@every 5s", func() {

// 			log.Info("del one dat old files in minio start")

// 			consumerHumanDetectionService := del_minio_one_month_age.NewService()
// 			err := consumerHumanDetectionService.DelOneDayOldFiles()
// 			if err != nil {
// 				log.Error(err)
// 			}

// 			log.Info("del one dat old files in minio done")
// 			fmt.Println("done")
// 		})

// 		cron.Start()
// 	}()

// 	wg.Wait()

// }
