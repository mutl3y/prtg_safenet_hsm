package safenet

//
//func PrtgChan(addr []string, count, size int, timeout, interval time.Duration, statsType string) error {
//	r := new(prtg.SensorResponse)
//	channels := make(chan prtg.SensorChannel, len(addr)*20)
//	errCH := make(chan error, 10)
//	wg := sync.WaitGroup{}
//
//	do := func(addr string, c chan prtg.SensorChannel, e chan error) {
//		addr = strings.TrimSpace(addr)
//		s, err := Ping(addr, count, size, timeout, interval)
//		if err != nil {
//			e <- err
//			if Debug {
//				fmt.Printf("Packet Loss:\t %+v\n", s)
//			}
//			wg.Done()
//			return
//		}
//
//		sch := prtg.SensorChannel{}
//		switch statsType {
//		case "loss":
//			sch = prtg.SensorChannel{
//				Name:      fmt.Sprintf("%v", addr),
//				Value:     s.PacketLoss,
//				Float:     1,
//				ShowChart: show,
//				ShowTable: show,
//				Unit:      prtg.UnitPercent,
//			}
//			if Debug {
//				fmt.Printf("ping response:\t %+v\n", *s)
//			}
//		default:
//			sch = prtg.SensorChannel{
//				Name:      fmt.Sprintf("%v", "avg Rtt"),
//				Value:     s.AvgRtt.Truncate(time.Microsecond).Seconds() * 1000,
//				Float:     1,
//				ShowChart: show,
//				ShowTable: show,
//				Unit:      prtg.UnitTimeResponse,
//			}
//		}
//
//		if statsType == "everything" {
//			scs := make([]prtg.SensorChannel, 0, 10)
//			scs = append(scs, prtg.SensorChannel{
//				Name:      fmt.Sprintf("%v", addr),
//				Value:     s.AvgRtt.Truncate(time.Microsecond).Seconds() * 1000,
//				Float:     1,
//				ShowChart: show,
//				ShowTable: show,
//				Unit:      prtg.UnitTimeResponse,
//			})
//			scs = append(scs, prtg.SensorChannel{
//				Name:  fmt.Sprintf("%v", "Packet Loss"),
//				Value: s.PacketLoss,
//				Float: 1,
//				Unit:  prtg.UnitPercent,
//			})
//			scs = append(scs, prtg.SensorChannel{
//				Name:  fmt.Sprintf("%v", "Packets Recv"),
//				Value: float64(s.PacketsRecv),
//				Unit:  prtg.UnitCount,
//			})
//			scs = append(scs, prtg.SensorChannel{
//				Name:  fmt.Sprintf("%v", "Packets Sent"),
//				Value: float64(s.PacketsSent),
//				Unit:  prtg.UnitCount,
//			})
//			scs = append(scs, prtg.SensorChannel{
//				Name:  fmt.Sprintf("%v", "MinRtt"),
//				Value: s.MinRtt.Round(time.Microsecond).Seconds() * 1000,
//				Float: 1,
//				Unit:  prtg.UnitTimeResponse,
//			})
//			scs = append(scs, prtg.SensorChannel{
//				Name:  fmt.Sprintf("%v", "MaxRtt"),
//				Value: s.MaxRtt.Round(time.Microsecond).Seconds() * 1000,
//				Float: 1,
//				Unit:  prtg.UnitTimeResponse,
//			})
//			scs = append(scs, prtg.SensorChannel{
//				Name:  fmt.Sprintf("%v", "StdDev Rtt"),
//				Value: s.StdDevRtt.Round(time.Microsecond).Seconds() * 1000,
//				Float: 1,
//				Unit:  prtg.UnitTimeResponse,
//			})
//
//			for _, s := range scs {
//				if Debug {
//					fmt.Println("sending stat")
//				}
//				c <- s
//			}
//
//		} else {
//			c <- sch
//		}
//		wg.Done()
//	}
//
//	if statsType == "everything" {
//		wg.Add(1)
//		do(addr[0], channels, errCH)
//	} else {
//		for _, v := range addr {
//			wg.Add(1)
//			go do(v, channels, errCH)
//		}
//	}
//
//	wg.Wait()
//	close(channels)
//	close(errCH)
//
//	for res := range channels {
//		r.AddChannel(res)
//	}
//
//	if len(errCH) >= 1 {
//		err := <-errCH
//		r.SensorResult.Error = 1
//		r.SensorResult.Text = fmt.Sprintf("%v", err)
//		fmt.Println(r.String())
//		return err
//	}
//
//	if Debug {
//		fmt.Println(r.Text)
//	} else {
//		fmt.Println(r.String())
//	}
//	return nil
//}
