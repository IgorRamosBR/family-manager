import { Button, Divider, Modal, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import React, { useState } from 'react';
import TransactionList from './list-transactions/list-transactions';
import AddTransactionForm from './add-transactions/add-transaction-form';
import { useQuery } from 'react-query';
import { CategoryApi } from '../../api/category';
import { ReportApi } from '../../api/report';
import { CategoryReportModel, ReportModel } from '../../models/report';
import { CategoryModel } from '../../models/category';

const Budget: React.FC = () => {
    const [isTransactionModalOpen, setIsTransactionModalOpen] = useState(false);

    // const { data: transactions } = useQuery(['report-transactions'], ReportApi.getReport);
    // const { data: categories } = useQuery(["categories-list"], CategoryApi.getCategories);

    const transactions:CategoryReportModel[] = []
    const categories:CategoryModel[] = []

    const showTransactionModal = () => {
        setIsTransactionModalOpen(true);
    };

    const handleOk = () => {
        setIsTransactionModalOpen(false);
    };

    const handleCancel = () => {
        setIsTransactionModalOpen(false);
    }; 


    return (
        <div>
            <div>
                <Space>
                    <Button onClick={showTransactionModal}>
                        <PlusOutlined /> Adicionar Transação
                    </Button>
                </Space>
                <Divider />

                <TransactionList transactions={transactions}/>

                <Modal title="Nova transação" open={isTransactionModalOpen} onOk={handleOk} onCancel={handleCancel}>
                    <AddTransactionForm categories={categories ? categories : []}/>
                </Modal>
            </div>
        </div>
    )
}

export default Budget;