/*

Package stream provides methods for listening to Twitter's Streaming APIs.

To begin using you will need credentials for an app setup in Twitter. Then you can get your timeline like so,

    // assuming consumerKey, consumerSecret, accessToken and accessSecret hold
    // the credentials
    auth := stream.Auth(consumerKey, consumerSecret, accessToken, accessSecret)

    for tweet := range.Timeline(auth) {
      log.Println(tweet)
    }

Warning, this library does not do any error handling at the moment, it ignores
any incoming messages that do not look like tweets.

*/
package stream
