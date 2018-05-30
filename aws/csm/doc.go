// Client Side Monitoring (CSM) is used to send metrics via UDP connection.
// Using the Start function will enable the reporting of metrics on a given
// port. If Start is not called prior to calling InjectHandler, then no
// handlers will be injected. That function will simply return the same
// handlers that were passed in.
//
// Pause can be called to pause any metrics publishing on a given port. Previous
// sessions that have had their handlers sessions modified via InjectHandlers
// may still be used. However, the handlers will act as a no-op meaning no metrics
// will be published.
//
//	Example:
//		r, err := csm.Start("clientID", ":31000")
//		if err != nil {
//			panic(fmt.Errorf("expected no error, but received %v", err))
//		}
//
//		sess := session.NewSession(&aws.Config{})
//		r.InjectHandlers(&sess.Handlers)
//
//		client := s3.New(sess)
//		resp, err := client.GetObject(&s3.GetObjectInput{
//			Bucket: aws.String("bucket"),
//			Key: aws.String("key"),
//		})
//
//		// Will pause monitoring
//		r.Pause()
//		resp, err = client.GetObject(&s3.GetObjectInput{
//			Bucket: aws.String("bucket"),
//			Key: aws.String("key"),
//		})
//
//		// Resume monitoring
//		r.Continue()
package csm
