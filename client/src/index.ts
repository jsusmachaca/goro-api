import express from 'express'
import path from 'node:path'
import { client } from './controllers/clientController'

const app = express()

app.use(express.static(path.join('src', 'public')))
app.set('view engine', 'ejs')
app.set('views', path.join(process.cwd(), 'src', 'views'))

app.get('/', client)

const PORT = process.env.PORT!
app.listen(PORT, () => {
  console.log(`Server listen on http://localhost:${PORT}`)
})
