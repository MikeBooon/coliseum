export function buildTestDomain(path: string, sub = 'test') {
    return `http://${sub}.localhost:6665${path}`
}
