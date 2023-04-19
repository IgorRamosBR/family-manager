import React from "react";
import { NavLink } from "react-router-dom";

export const NavBarBrand: React.FC = () => {
  return (
    <div className="nav-bar__brand">
        <NavLink to="/">
          <h1 id="page-title" className="content__title">
            Controle Familiar
          </h1>
        </NavLink>
      {/* <NavLink to="/">
        <img
          className="nav-bar__logo"
          src="https://cdn.auth0.com/blog/hub/code-samples/hello-world/auth0-logo.svg"
          alt="Auth0 shield logo"
          width="122"
          height="36"
        />
      </NavLink> */}
    </div>
  );
};
