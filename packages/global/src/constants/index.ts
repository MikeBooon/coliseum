export const INPUTS = {
    MAX_EMAIL_LENGTH: 300,
    MAX_NAME_LENGTH: 200,
    MAX_SLUG_LENGTH: 100,
    MIN_NAME_LENGTH: 2,
    MIN_SLUG_LENGTH: 2,
    PASSWORD_MIN_LENGTH: 10,
    PASSWORD_REGEX:
        '^.*(?=.{8,})(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+\\-=\\[\\]{};\':"\\\\|,.<>\\/?]).*$',
} as const
