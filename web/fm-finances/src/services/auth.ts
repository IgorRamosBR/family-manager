function getToken(): string {
    const token = sessionStorage.getItem('token')
    const tokenTime = sessionStorage.getItem('tokenTime')

    if (!token) {
        return ''
    }

    if (tokenTime) {
        const ttl = new Date().getTime() - parseInt(tokenTime)
        if (ttl > 3600000) {
            return ""
        }

        return token
    }

    return ''
}

export const Auth = {
    getToken
}