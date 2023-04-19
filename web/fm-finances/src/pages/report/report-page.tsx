import React from 'react';
import ReportList from './report-list';
import { PageLayout } from 'src/components/page-layout';
import { ReportApi } from 'src/services/report.service';
import { useQuery } from 'react-query';
import { PageLoader } from 'src/components/page-loader';

const ReportPage: React.FC = () => {
    const { data: transactions, isLoading } = useQuery(['report-transactions'], ReportApi.getReport);


    if (isLoading) {
        return (
            <div className="page-layout">
                <PageLoader />
            </div>
        );
    }

    return (
        <PageLayout>
            <>
                <div className="content-layout">
                    <h1 id="page-title" className="content__title">
                        Relatório Financeiro de {new Date().getFullYear()}
                    </h1>
                    <div className="content__body">
                        <ReportList transactions={transactions} />
                    </div>
                </div>
            </>
        </PageLayout>
    )
}

export default ReportPage