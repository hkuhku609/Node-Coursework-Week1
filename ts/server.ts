// server.js
// This is where your node app starts

//load the 'express' module which makes writing webservers easy
require('dotenv').config()

import cors from 'cors'
import express, { NextFunction, Request, Response } from 'express'
import { AddressInfo } from 'net'
import quotes from './routes/quotes'

const app = express()
app.use(cors())
app.use(requestLogger)
// Now register handlers for some routes:
//   /                  - Return some helpful welcome info (text)
//   /quotes            - Should return all quotes (json)
//   /quotes/random     - Should return ONE quote (json)
app.get('/', function (request: Request, response: Response) {
  response.send("Neill's Quote Server!  Ask me for /quotes/random, or /quotes")
})

app.use('/quotes', quotes)
//Start our server so that it listens for HTTP requests!
const listener = app.listen(process.env.PORT || 4000, () => {
  const { port } = listener.address() as AddressInfo
  console.log('Your app is listening on port ' + port)
})

function requestLogger(req: Request, res: Response, next: NextFunction) {
  const startTime = process.hrtime()

  res.on('finish', () => {
    const elapsedTime = process.hrtime(startTime)
    const elapsedTimeInMs = elapsedTime[0] * 1000 + elapsedTime[1] / 1e6
    console.log(
      `[${req.method}] [${req.url}] [${elapsedTimeInMs.toFixed(3)}ms]`
    )
  })

  next()
}
//
