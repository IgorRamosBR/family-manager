import { Table } from "antd";
import type { ColumnsType } from "antd/es/table";
import React, { useEffect, useState } from "react";
import { CategoryReportModel } from "../../../models/report";
import { monthCodes, monthNames } from "../../../utils/utils";

export interface DataType {
    key: string;
    category: string;
    month1?: number;
    month2?: number;
    month3?: number;
    month4?: number;
    month5?: number;
    month6?: number;
    total: number;
    isCategory: boolean
}

const columns: ColumnsType<DataType> = [
    {
        title: "",
        key: "category",
        dataIndex: "category",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{record.category}</span>)
            }
            return (<span>{record.category}</span>)
        }
    },
    {
        title: getMonth(0),
        key: "month1",
        dataIndex: "month1",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(Number(record.month1))}</span>)
            }
            return (<span>{new Intl.NumberFormat("pt-BR").format(Number(record.month1))}</span>)
        }
    },
    {
        title: getMonth(1),
        key: "month2",
        dataIndex: "month2",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(Number(record.month2))}</span>)
                }
            return (<span>{new Intl.NumberFormat("pt-BR").format(Number(record.month2))}</span>)
        }
    },
    {
        title: getMonth(2),
        key: "month3",
        dataIndex: "month3",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(Number(record.month3))}</span>)
            }
            return (<span>{new Intl.NumberFormat("pt-BR").format(Number(record.month3))}</span>)
        }
    },
    {
        title: getMonth(3),
        key: "month4",
        dataIndex: "month4",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(Number(record.month4))}</span>)
            }
            return (<span>{new Intl.NumberFormat("pt-BR").format(Number(record.month4))}</span>)
        }
    },
    {
        title: getMonth(4),
        key: "month5",
        dataIndex: "month5",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(Number(record.month5))}</span>)
            }
            return (<span>{new Intl.NumberFormat("pt-BR").format(Number(record.month5))}</span>)
        }
    },
    {
        title: getMonth(5),
        key: "month6",
        dataIndex: "month6",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(Number(record.month6))}</span>)
            }
            return (<span>{new Intl.NumberFormat("pt-BR").format(Number(record.month6))}</span>)
        }
    },
    {
        title: "Total",
        key: "total",
        dataIndex: "total",
        render: (_, record) => {
            if (record.isCategory) {
                return (<span style={{ fontWeight: "bold", textTransform: "uppercase" }}>{new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(Number(record.total))}</span>)
            }
            return (<span>{new Intl.NumberFormat("pt-BR").format(Number(record.total))}</span>)
        }
    }
]

function getMonth(index: number) {
    const date = new Date();
    let month = date.getMonth();
    if (month <= 6) {
        return monthNames[index]
    }
    return monthNames[index + 5]
}

function getMonthCode(index: number) {
    const date = new Date();
    let month = date.getMonth();
    if (month <= 6) {
        return monthCodes[index]
    }
    return monthCodes[index + 5]
}

function getMonthValue(index: number, values: Map<string, number>): number {
    let monthCode = getMonthCode(index)
    let map = new Map(Object.entries(values))
    let value = map.get(monthCode)
    return value ? value : 0.0
}
interface Props {
    transactions: CategoryReportModel[] | undefined
}

function TransactionList({ transactions }: Props) {
    const [rows, setRows] = useState<DataType[]>([])

    

    useEffect(() => {
        let newRows: DataType[] = []
        transactions?.forEach(t => {
            if (t.isParent) {
                let row: DataType = {
                    key: t.name,
                    category: t.name,
                    total: t.total,
                    isCategory: true
                }
                newRows.push(row)
                return
            }
            let row: DataType = {
                key: t.name,
                category: t.name,
                total: t.total,
                month1: getMonthValue(0, t.values),
                month2: getMonthValue(1, t.values),
                month3: getMonthValue(2, t.values),
                month4: getMonthValue(3, t.values),
                month5: getMonthValue(4, t.values),
                month6: getMonthValue(5, t.values),
                isCategory: false
            }
            newRows.push(row)
        })
        setRows(newRows)
    }, [transactions])

    return (
        <Table columns={columns} dataSource={rows} pagination={false}/>
    )
}

export default TransactionList;