import { Button, Divider, Modal, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import React, { useState } from 'react';
import TransactionForm from 'src/components/transaction-form';
import { PageLayout } from 'src/components/page-layout';
import { TransactionApi } from 'src/services/transaction.service';
import { useQuery } from 'react-query';
import TransactionList from './transaction-list';
import { PageLoader } from 'src/components/page-loader';
import { CategoryApi } from 'src/services/category.service';

const TransactionsPage: React.FC = () => {
    const [isTransactionModalOpen, setIsTransactionModalOpen] = useState(false);
    const {data: transactions,isLoading: transactionsLoading, refetch } = useQuery("transactions-list", TransactionApi.getLatestTransactions)
    const {data: categories, isLoading: categoriesLoading } = useQuery("categories-list", CategoryApi.getCategories)

    const showTransactionModal = () => {
        setIsTransactionModalOpen(true);
    };

    const handleOk = () => {
        setIsTransactionModalOpen(false);
    };

    const handleClose = () => {
        refetch()
        setIsTransactionModalOpen(false);
    };

    if (transactionsLoading || categoriesLoading) {
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
                        Transações
                    </h1>
                    <div className="content__body">
                        <Space>
                            <Button type='text' onClick={showTransactionModal}>
                                <PlusOutlined /> Adicionar Transação
                            </Button>
                        </Space>
                        <Divider />

                        <TransactionList transactions={transactions}/>

                        <Modal title="Nova transação" open={isTransactionModalOpen} footer={null} onOk={handleOk} onCancel={handleClose}>
                            <TransactionForm categories={categories ? categories : []} onFinish={handleClose} />
                        </Modal>
                    </div>
                </div>
            </>
        </PageLayout>

    )
}

export default TransactionsPage;