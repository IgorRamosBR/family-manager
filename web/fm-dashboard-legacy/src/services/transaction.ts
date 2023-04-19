import { resolve } from "path";
import { TransactionApi } from "../api/transaction";
import { TransactionModel } from "../models/transaction";
import { monthNames } from "../utils/utils";

interface TransactionType {
    key: string;
    category: string;
    month1: number;
    month2: number;
    month3: number;
    month4: number;
    month5: number;
    month6: number;
    total: number;
    isHeader: boolean
}

interface CategoryType {
    name: string;
    type: string;
}

// async function getAllTransactionsOrdered(): TransactionType[] {
//     let transactionsByCategory = new Map<string, TransactionType>()
//     let categories = await initCategories()
//     let periods = getPeriods()
//     periods.forEach((period, index) => {
//         TransactionRepository.getTransactions(period)
//             .then(transactions => {
//                 transactions.forEach(t => {
//                     let newTransaction = mergeTransactionType(transactionsByCategory.get(t.category), index, t)
//                     transactionsByCategory.set(t.category, newTransaction)
//                 })
//             })
//     })


// }

async function initCategories() {
    const categories: CategoryType[] = [
        {
            name: "Renda Familiar#Salário Igor",
            type: "INCOME",
        },
        {
            name: "Renda Familiar#Salário Camilla",
            type: "INCOME",
        },
        {
            name: "Habitação#Apartamento",
            type: "EXPENSE",
        }
    ]

    let transactions = new Map()
    categories.forEach(function (item) {
        transactions.set(item.name, item.type)
    })
}

function getPeriods() {
    let periods = new Array()
    let months = new Array()
    const date = new Date();
    const month = date.getMonth();
    const year = date.getFullYear()
    if (month <= 6) {
        months = monthNames.slice(0, 6)
    }
    months = monthNames.slice(6, 12)
    months.forEach(function (item) {
        periods.push(`${item.substring(0, 3)}-${year}`)
    })

    return periods
}



// export const transactionService = {
//     getAllTransactionsOrdered
// }
function mergeTransactionType(transaction: TransactionType | undefined, month: number, transactionModel: TransactionModel) {
    if (transaction) {
       return updateTransactionValue(transaction, month, transactionModel.value)
    }
    let newTransaction: TransactionType = {
        key: transactionModel.transactionId,
        category: transactionModel.category,
        month1: 0,
        month2: 0,
        month3: 0,
        month4: 0,
        month5: 0,
        month6: 0,
        total: 0,
        isHeader: false,

    }
    return updateTransactionValue(newTransaction, month, transactionModel.value)

    throw new Error("Function not implemented.");
}

function updateTransactionValue(transaction: TransactionType, index: number, value: number): TransactionType {
    switch (index) {
        case 0:
            transaction.month1 += value
            break
        case 1:
            transaction.month2 += value
            break
        case 2:
            transaction.month3 += value
            break
        case 3:
            transaction.month4 += value
            break
        case 4:
            transaction.month5 += value
            break
        case 5:
            transaction.month6 += value
            break
    }
    return transaction;
}
