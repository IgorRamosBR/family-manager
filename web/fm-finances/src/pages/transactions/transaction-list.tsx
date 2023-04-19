import React, { useEffect, useState } from 'react';
import { Table, Tag } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import { TransactionModel } from 'src/models/transaction';

interface DataType {
    key: React.Key;
    description: string;
    value: number;
    category: string;
    period: string;
    date: string;
    type: string;
    paymentMethod: string;
}

const columns: ColumnsType<DataType> = [
    {
        title: 'Categoria',
        dataIndex: 'category',
        key: 'category',
    },
    {
        title: 'Valor',
        dataIndex: 'value',
        key: 'value',
        render: (_, record) => {
            return (<span>{new Intl.NumberFormat("pt-BR", { style: 'currency', currency: 'BRL' }).format(Number(record.value))}</span>)
        }
    },
    {
        title: 'Descrição',
        dataIndex: 'description',
        key: 'description',
    },
    {
        title: 'Período',
        dataIndex: 'period',
        key: 'period',
    },
    {
        title: 'Data',
        dataIndex: 'date',
        key: 'date',
    },
    {
        title: 'Meio de pagamento',
        dataIndex: 'paymentMethod',
        key: 'paymentMethod',
        render: (_, record) => {
            return (
                <Tag color="blue">
                    {record.paymentMethod === 'CREDIT_CARD' ? 'CRÉDITO' : record.paymentMethod === 'DEBIT_CARD' ? 'DÉBITO' : record.paymentMethod
                    }
                </Tag>
            )
            
        }
    },
    {
        title: 'Tipo',
        dataIndex: 'type',
        key: 'type',
        render: (_, record) => {
            if (record.type === 'INCOME') {
                return (
                    <Tag color={'green'}>
                        RENDA
                    </Tag>
                )
            }
            return (
                <Tag color={'volcano'}>
                    DESPESA
                </Tag>
            )
        }
    },
];

interface Props {
    transactions: TransactionModel[] | undefined
}

const TransactionList: React.FC<Props> = ({ transactions }) => {
    const [transactionsData, setTransactionsData] = useState<DataType[]>([])

    const mapTransactions = (transactions: TransactionModel[]): DataType[] => {
        return transactions.map(t => {
            return {
                key: t.transactionId,
                description: t.description,
                value: t.value,
                category: t.category,
                period: t.monthYear,
                date: t.date,
                type: t.type,
                paymentMethod: t.paymentMethod
            } as DataType;
        });
    }

    useEffect(() => {
        if (transactions) {
            const transactionsDataType = mapTransactions(transactions);
            setTransactionsData(transactionsDataType)
        }


    }, [transactions])

    return (
        <Table columns={columns} dataSource={transactionsData} />
    )
}
export default TransactionList;