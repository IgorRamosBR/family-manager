import { useAuth0 } from "@auth0/auth0-react";
import React from "react";
import { PageLoader } from "./components/page-loader";
import { AuthenticationGuard } from "./components/authentication-guard";
import { Route, Routes } from "react-router-dom";
import { AdminPage } from "./pages/admin-page";
import { CallbackPage } from "./pages/callback-page";
import { NotFoundPage } from "./pages/not-found-page";
import TransactionsPage from "./pages/transactions/transaction-page";
import ReportPage from "./pages/report/report-page";
import { QueryClient, QueryClientProvider } from "react-query";
import CategoriesPage from "./pages/categories/categories-page";

const queryClient = new QueryClient();

export const App: React.FC = () => {
  const { isLoading } = useAuth0();

  if (isLoading) {
    return (
      <div className="page-layout">
        <PageLoader />
      </div>
    );
  }
  return (
    <QueryClientProvider client={queryClient}>
      <Routes>
        <Route path="/" element={<ReportPage />} />
        <Route
          path="/report"
          element={<AuthenticationGuard component={ReportPage} />}
        />
        <Route path="/transactions" element={<TransactionsPage />} />
        <Route
          path="/categories"
          element={<AuthenticationGuard component={CategoriesPage} />}
        />
        <Route
          path="/admin"
          element={<AuthenticationGuard component={AdminPage} />}
        />
        <Route path="/callback" element={<CallbackPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </QueryClientProvider>
  );
};
