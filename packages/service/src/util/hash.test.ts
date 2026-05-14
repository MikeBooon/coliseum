import { describe, expect, it } from 'vitest'
import { compareHash, createHash } from './hash.ts'

const TEST_PASSWORDS = ['hunter2', '%213n1(@ 129j12,n', 'ñöñ-äśçïï', 'password🚀123']

describe('password hashing', async () => {
    it.for(TEST_PASSWORDS)('should hash password and then verify it', async (password) => {
        const hash = await createHash(password)

        expect(hash).toBeDefined()

        expect(await compareHash('wrong', hash)).toBe(false)
        expect(await compareHash('', hash)).toBe(false)
        expect(await compareHash(password, hash)).toBe(true)
        expect(await compareHash('-', hash)).toBe(false)
        expect(await compareHash(password, hash)).toBe(true)
    })
})
