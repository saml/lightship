package saml.lightship

import io.undertow.Undertow
import io.undertow.server.{HttpServerExchange, HttpHandler}

/**
 * Created by saml on 3/9/15.
 */
object Main extends App {
  val handler = new HttpHandler {
    override def handleRequest(httpServerExchange: HttpServerExchange): Unit =
      httpServerExchange.getResponseSender.send("Hello world")
  }
  val host = "localhost"
  val port = 8080
  val server = Undertow.builder()
    .addHttpListener(port, host)
    .setHandler(handler)
    .build()

  server.start()
  println(s"Listening to http://${host}:${port}")

  val console = System.console()
  if (console != null) {
    console.readLine("Press enter to exit")
    server.stop()
  }
}
