import { CategoryModel } from "./category";

export interface TransactionModel {
    transactionId: string;
    description: string;
    monthYear: string;
    value: number;
    paymentMethod: string,
    type: string;
    category: string;
    date: string;
}

export interface TransactionPageModel {
    results: TransactionModel[],
    next: string,
}