import { useAuth0 } from "@auth0/auth0-react";
import React, { useEffect } from "react";
import { Navigate } from "react-router-dom";
import { NavBar } from "../components/navigation/desktop/nav-bar";
import { MobileNavBar } from "../components/navigation/mobile/mobile-nav-bar";
import { PageLayout } from "../components/page-layout";

export const CallbackPage: React.FC = () => {
  const { error, getAccessTokenSilently } = useAuth0();

  useEffect(() => {
    try {
      const getToken = async () => {
        const accessToken = await getAccessTokenSilently()
        sessionStorage.setItem('token', accessToken)
        sessionStorage.setItem('tokenTime', new Date().getTime().toString())
      }

      getToken()
    } catch (e) {
      console.log('error to get token')
    }
  })

  if (error) {
    return (
      <PageLayout>
        <div className="content-layout">
          <h1 id="page-title" className="content__title">
            Error
          </h1>
          <div className="content__body">
            <p id="page-description">
              <span>{error.message}</span>
            </p>
          </div>
        </div>
      </PageLayout>
    );
  }



  return (
    <div className="page-layout">
      <NavBar />
      <MobileNavBar />
      <Navigate to={"/report"} />
      <div className="page-layout__content" />
    </div>
  );
};
