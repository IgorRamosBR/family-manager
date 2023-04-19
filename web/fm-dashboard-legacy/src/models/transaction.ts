export interface TransactionModel {
    transactionId: string;
    description: string;
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