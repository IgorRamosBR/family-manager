import { useAuth0 } from "@auth0/auth0-react";
import { ConfigProvider, theme } from "antd";
import React from "react";
import { NavBar } from "./navigation/desktop/nav-bar";
import { MobileNavBar } from "./navigation/mobile/mobile-nav-bar";
import { PageFooter } from "./page-footer";

interface Props {
  children: JSX.Element;
}

const { darkAlgorithm } = theme;

export const PageLayout: React.FC<Props> = ({ children }) => {
  const { isAuthenticated, loginWithRedirect } = useAuth0();

  if (!isAuthenticated) {
    loginWithRedirect({
      appState: {
        returnTo: "/callback",
      },
      authorizationParams: {
        prompt: "login",
      },
    });
  }

  return (
    <ConfigProvider theme={{ algorithm: darkAlgorithm }}>
      <div className="page-layout">
        <NavBar />
        <MobileNavBar />
        <div className="page-layout__content">{children}</div>
        <PageFooter />
      </div>
    </ConfigProvider>
  );
};
