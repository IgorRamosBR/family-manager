import { ConfigProvider, theme } from "antd";
import React from "react";
import { Auth0Features } from "src/components/auth0-features";
import { HeroBanner } from "src/components/hero-banner";
import { PageLayout } from "../components/page-layout";
import ReportPage from "./report/report-page";

const { darkAlgorithm } = theme;

export const HomePage: React.FC = () => (
  <PageLayout>
    <>

      <ReportPage />

      {/* <HeroBanner />
      <Auth0Features /> */}
    </>
  </PageLayout>
);
