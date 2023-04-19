import { useAuth0 } from "@auth0/auth0-react";
import React, { useEffect } from "react";
import { useHistory } from "react-router-dom";

export function Callback() {
    const { user, isAuthenticated, getAccessTokenSilently } = useAuth0();
    const history = useHistory()

    useEffect(() => {
        console.log('passei aqui')
        try {
            const getToken = async () => {
                const accessToken = await getAccessTokenSilently()
                localStorage.setItem('token', accessToken)
                localStorage.setItem('tokenTime', new Date().getTime().toString())
            }
            
            getToken()
        } catch (e) {
            console.log('error to get token')
        }
       
    },[getAccessTokenSilently, user])

    useEffect(() => {
        console.log(isAuthenticated)
    })

    return (
        <>
         {history.push("/")}
        </>
    )
}