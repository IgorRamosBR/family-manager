export interface TransactionModel {
    transactionId: string;
    monthYear: string;
    value: number;
    type: string;
    category: string;
    date: string;
}

export interface TransactionPageModel {
    results: TransactionModel[],
    next: string,
}