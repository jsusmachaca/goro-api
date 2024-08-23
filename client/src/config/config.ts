import jwt from 'jsonwebtoken'
import fs from 'node:fs/promises'
import path from 'node:path'

export const genJWT = async () => {
  const privKey = await fs.readFile(path.join(process.cwd(), 'jwt.key'))
  const token = jwt.sign(
    {
      id: "79da5bdb-de55-43f4-b29f-50393050a9c0"
    }, 
    privKey, { algorithm: 'RS256' }
  )
  return token
}