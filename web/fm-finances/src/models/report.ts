export interface ReportModel {
    categories: Map<string, CategoryReportModel>;
}

export interface CategoryReportModel {
    name: string;
    total: number;
    isParent: boolean;
    values: Map<string, number>;
}