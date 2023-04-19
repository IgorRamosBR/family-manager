import { useAuth0 } from "@auth0/auth0-react";
import React from "react";
import { NavBarTab } from "./nav-bar-tab";

export const NavBarTabs: React.FC = () => {
  const { isAuthenticated } = useAuth0();

  return (
    <div className="nav-bar__tabs">
      {isAuthenticated && (
        <>
          <NavBarTab path="/report" label="Finanças" />
          <NavBarTab path="/transactions" label="Transaçooes" />
          <NavBarTab path="/categories" label="Categorias" />
        </>
      )}
    </div>
  );
};
