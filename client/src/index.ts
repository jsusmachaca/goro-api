// process.loadEnvFile() // uncomment in development mode
import express from 'express'
import path from 'node:path'
import axios from 'axios'
import { RequestData } from './types/RequestData'

const app = express()

const API_SERVER = process.env.API_SERVER!

app.use(express.static(path.join('src', 'public')))
app.set('view engine', 'ejs')
app.set('views', path.join(process.cwd(), 'src', 'views'))

app.get('/', async (req, res) => {
  const { name } = req.query
  let pageRes: RequestData[] = []

  for (let i = 1; i < 3; i++) {
    const reqServer = await axios.get(`http://${API_SERVER}/api/RnM?page=${i}`)
    pageRes = pageRes.concat(reqServer.data)
  }

  if (name && name !== '') {
    const filterData = pageRes.filter(
      (data: RequestData) => data.name.toLowerCase().startsWith((name as string).toLowerCase())
    )
    return res.render('index', { data: filterData })
  }

  return res.render('index', { data: pageRes })
})

const PORT = process.env.PORT!
app.listen(PORT, () => {
  console.log(`Server listen on http://localhost:${PORT}`)
})