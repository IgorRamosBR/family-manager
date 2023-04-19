import { TransactionModel, TransactionPageModel } from "../models/transaction";
import { monthCodes } from "../utils/utils";
import { Auth } from "./auth";
import dayjs from 'dayjs';

const URL = `${process.env.REACT_APP_API_SERVER_URL}/transactions`;
const defaultOptions = {
    headers: {
        'Authorization': Auth.getToken(),
        'Content-Type': 'application/json',
    },
};

async function getTransactions(period: string, offset: string): Promise<TransactionPageModel> {
    return fetch(`${URL}?period=${period}&offset=${offset}`, {
        method: 'GET',
        ...defaultOptions
    })
        .then(response => response.json())
        .then(response => { return response as TransactionPageModel })
}

async function createTransaction(transaction: TransactionModel) {
    if (!transaction.date) {
        transaction.date = dayjs().format('DD/MMM/YYYY')
    }
    transaction.monthYear = `${monthCodes[dayjs().month()]}-${dayjs().year()}`
    let response = await fetch(URL, {
        method: 'POST',
        body: JSON.stringify(transaction),
        ...defaultOptions
    })

    if (!response?.ok) {
        const responseBody = await response.json()
        throw new Error(responseBody)
    }
}

async function getAllTransactionsByPeriod(period: string, offset: string): Promise<TransactionModel[]> {
    const page = await getTransactions(period, offset);
    let transactions: TransactionModel[] = []
    transactions.push(...page.results)

    if (page.next) {
        const newResponse = await getAllTransactionsByPeriod(period, page.next)
        transactions.push(...newResponse)
    }

    return transactions
}

async function getLatestTransactions(): Promise<TransactionModel[]> {
    const year = new Date().getFullYear();
    const months = getLatestMonthCodes();

    let transactions = getMonthlyTransactions(months, year)
    return transactions
}

async function getSixMonthsTransactions(): Promise<TransactionModel[]> {
    console.log('passei aqui')
    const year = new Date().getFullYear();
    const months = getCurrentSixMonths();

    let transactions = getMonthlyTransactions(months, year)
    return transactions
}

async function getMonthlyTransactions(months: string[], year: number): Promise<TransactionModel[]> {
    let transactions: TransactionModel[] = []
    for (let i = 0; i < months.length; i++) {
        let transactionsByPeriod = await getAllTransactionsByPeriod(`${months[i]}-${year}`, "")
        transactions.push(...transactionsByPeriod)
    }

    transactions.sort((d1, d2) => parseDate(d2.date).getTime() - parseDate(d1.date).getTime())
    
    return transactions
}

const parseDate = (dateString: string): Date => {
    const parts = dateString.split('/');
    const day = parseInt(parts[0]);
    const month = parseInt(parts[1]) - 1;
    const year = parseInt(parts[2]);
    return new Date(year, month, day);
  };

  
const getCurrentSixMonths = (): string[] => {
    const date = new Date();
    let month = date.getMonth();
    if (month < 6) {
        return monthCodes.slice(0, 6)
    }
    return monthCodes.slice(7, monthCodes.length)
}

const getLatestMonthCodes = (): string[] => {
    const date = new Date();
    let month = date.getMonth();
    return monthCodes.slice(0, month + 1)
}

export const TransactionApi = {
    getTransactions,
    createTransaction,
    getSixMonthsTransactions,
    getLatestTransactions
}

