import { scrypt, timingSafeEqual } from 'crypto'
import { nanoid } from 'nanoid'

const HASH_KEY_LENGTH = 64
const SALT_LENGTH = 24

const HASH_FORMAT = 'base64url'

const SEPARATOR = '/'

const CURRENT_VERSION = '1'

export async function createHash(value: string): Promise<string> {
    return new Promise((resolve, reject) => {
        const salt = nanoid(SALT_LENGTH)
        scrypt(value, salt, HASH_KEY_LENGTH, (err, derivedKey) => {
            if (err) {
                return reject(err)
            }

            const key = derivedKey.toString(HASH_FORMAT)

            resolve(`${CURRENT_VERSION}${SEPARATOR}${salt}${SEPARATOR}${key}`)
        })
    })
}

export async function compareHash(value: string, hash: string): Promise<boolean> {
    return new Promise((resolve, reject) => {
        const [version, salt, originKeyB64] = hash.split(SEPARATOR)
        if (!version || !salt || !originKeyB64) {
            throw new Error('Invalid hash')
        }

        const originalKey = Buffer.from(originKeyB64, HASH_FORMAT)

        scrypt(value, salt, HASH_KEY_LENGTH, (err, derivedKey) => {
            if (err) {
                return reject(err)
            }

            resolve(timingSafeEqual(originalKey, derivedKey))
        })
    })
}
