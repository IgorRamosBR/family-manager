import { Table, message } from "antd";
import { ColumnsType } from "antd/es/table";
import React, { useEffect, useState } from "react";
import { TransactionApi } from '../../api/transaction';
import { TransactionPageModel } from "../../models/transaction";
import { getCurrentPeriod } from "../../utils/utils";

interface TransactionDataType {
    key: string;
    period: string;
    description: string;
    category: string;
    value: number;
    date: string;
}

const columns: ColumnsType<TransactionDataType> = [
    {
        title: 'Período',
        dataIndex: 'period',
        key: 'period',
    },
    {
        title: 'Descrição',
        dataIndex: 'description',
        key: 'description',
    },
    {
        title: 'Categoria',
        dataIndex: 'period',
        key: 'period',
    },
    {
        title: 'Valor',
        dataIndex: 'value',
        key: 'value',
    },
    {
        title: 'Data',
        dataIndex: 'date',
        key: 'date',
    }
]


function TransactionsPage() {
    const [transactions, setTransactions] = useState<TransactionDataType[]>([])
    const [offset, setOffset] = useState<string>('')
    const [messageApi, contextHolder] = message.useMessage();

    useEffect(() => {
        const fetchTransactions = async () => {
            let transactions = await TransactionApi.getTransactions(getCurrentPeriod(), offset)
            console.log(transactions)
            setTransactions(mapToTransactionType(transactions))
            console.log('passei aqui')
        }
        try {
            fetchTransactions()
        } catch (e) {
            messageApi.open({
                type: 'error',
                content: 'Erro ao buscar transações',
            });
        }
    })

    const mapToTransactionType = (transactionsPage: TransactionPageModel) => {
        let transactionsData = []
        const results = transactionsPage.results
        for (let i = 0; i < results.length; i++) {
            const transactionType: TransactionDataType = {
                key: i.toString(),
                period: results[i].monthYear,
                description: results[i].description,
                category: results[i].category,
                value: results[i].value,
                date: results[i].date
            }
            transactionsData.push(transactionType)
        }
        return transactionsData
    }

    return (
        <>
            {contextHolder}
            <Table columns={columns} dataSource={transactions} pagination={false} />
        </>
    )
}

export default TransactionsPage;