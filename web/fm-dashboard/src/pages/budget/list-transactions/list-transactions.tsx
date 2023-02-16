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
    },
    {
        title: getMonth(1),
        key: "month2",
        dataIndex: "month2",
    },
    {
        title: getMonth(2),
        key: "month3",
        dataIndex: "month3",
    },
    {
        title: getMonth(3),
        key: "month4",
        dataIndex: "month4",
    },
    {
        title: getMonth(4),
        key: "month5",
        dataIndex: "month5",
    },
    {
        title: getMonth(5),
        key: "month6",
        dataIndex: "month6",
    },
    {
        title: "Total",
        key: "total",
        dataIndex: "total"
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