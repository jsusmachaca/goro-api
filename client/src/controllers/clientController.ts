process.loadEnvFile() // uncomment in development mode
import { Request, Response } from "express"
import axios from 'axios'
import { RequestData } from '../types/RequestData'
import { genJWT } from '../config/config'

const API_SERVER = process.env.API_SERVER!

export const client = async (req: Request, res: Response) => {
  try {
    const { name } = req.query
    let pageRes: RequestData[] = []
    const token = await genJWT()

    const reqServer: any  = await axios.get(`http://${API_SERVER}/api/rnm`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })

    pageRes = reqServer.data
    
    if (name && name !== '') {
      const filterData = pageRes.filter(
        (data: RequestData) => data.name.toLowerCase().startsWith((name as string).toLowerCase())
      )
      return res.render('index', { data: filterData })
    }

    return res.render('index', { data: pageRes })
  } catch (err) {
    return res.status(500).json({ error: "an error ocurred" })
  }
}
