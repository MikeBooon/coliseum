export abstract class RequestError extends Error {
    public id: string
    public statusCode: number
    constructor(statusCode: number, message: string) {
        super(message)
        this.statusCode = statusCode
        this.id = this.constructor.name
    }
}

export class UniqueConstraintError extends RequestError {
    public property: string
    constructor(property: string) {
        super(400, `unique constraint property: '${property}'`)
        this.property = property
    }
}
