import Foundation
import Kitura

let router = Router()

router.get("/ok") { request, response, next in
  response.send("ok")
}

Kitura.addHTTPServer(onPort: 8080, with: router)

Kitura.run()