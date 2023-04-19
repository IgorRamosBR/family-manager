import { CategoryReportModel } from "src/models/report";
import { monthCodes } from "src/utils/utils";
import { Auth } from "./auth";

const URL = `${process.env.REACT_APP_API_SERVER_URL}/report`;
const defaultOptions = {
    headers: {
        'Authorization': Auth.getToken(),
        'Content-Type': 'application/json',
    },
};

async function getReport(): Promise<CategoryReportModel[]> {
    let periods = getCurrentSixMonths()
    return fetch(`${URL}?periods=${periods.join(",")}`, {
        method: 'GET',
        ...defaultOptions
    })
        .then(response => response.json())
        .then(response => { return response as CategoryReportModel[] })
}

const getCurrentSixMonths = ():string[] => {   
    const date = new Date();
    let month = date.getMonth();
    if (month < 6) {
        return monthCodes.slice(0, 6).map(m => m + `-${date.getFullYear()}`)
    }
    return monthCodes.slice(7, monthCodes.length).map(m => m + `-${date.getFullYear()}`)
}

export const ReportApi = {
    getReport,
}