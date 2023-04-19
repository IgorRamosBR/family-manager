import { Auth } from "../components/auth/auth";
import { TransactionModel, TransactionPageModel } from "../models/transaction";
import { monthCodes } from "../utils/utils";

const URL = `${process.env.REACT_APP_FINANCES_API}/transactions`;
const defaultOptions = {
    headers: {
        'Authorization': Auth.getToken(),
        'Content-Type': 'application/json',
        "Access-Control-Allow-Origin":      "*",
    },
};

async function getTransactions(period:string, offset: string): Promise<TransactionPageModel> {
    return fetch(`${URL}?period=${period}&offset=${offset}`, {
        method: 'GET',
        ...defaultOptions
    })
        .then(response => response.json())
        .then(response => {return response as TransactionPageModel})
}

async function createTransaction(transaction: TransactionModel) {
    const response = await fetch(URL, {
        method: 'POST',
        body: JSON.stringify(transaction),
        ...defaultOptions
    })
}

async function getAllTransactionsByPeriod(period:string, offset: string): Promise<TransactionModel[]> {
    const page = await getTransactions(period, offset);
    let transactions: TransactionModel[] = []
    transactions.push(...page.results)

    if (page.next) {
        const newResponse = await getAllTransactionsByPeriod(period, page.next)
        transactions.push(...newResponse)
    }

    return transactions
}

async function getSixMonthsTransactions(): Promise<TransactionModel[]> {
    const year = new Date().getFullYear();
    const months = getCurrentSixMonths();

    let transactions: TransactionModel[] = []
    for (let i = 0; i < months.length; i++) {
        let transactionsByPeriod = await getAllTransactionsByPeriod(`${months[i]}-${year}`, "")
        transactions.push(...transactionsByPeriod)
    }
   
    return transactions
}

const getCurrentSixMonths = ():string[] => {   
    const date = new Date();
    let month = date.getMonth();
    if (month < 6) {
        return monthCodes.slice(0, 6)
    }
    return monthCodes.slice(7, monthCodes.length)
}

export const TransactionApi = {
    getTransactions,
    createTransaction,
    getSixMonthsTransactions
}