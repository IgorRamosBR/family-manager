import { DatePicker, Form, Input, InputNumber, Radio, Select } from 'antd';
import React from 'react';
import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import { CategoryModel } from '../../../models/category';

dayjs.extend(customParseFormat);
const dateFormatList = ['DD/MM/YYYY', 'DD/MM/YY'];

const getTodayDate = (): string => {
    const toString = (date: number, padLength: number) => {
        return date.toString().padStart(padLength, '0');
    };

    const date = new Date();

    const dateTimeNow = toString(date.getHours(), 2) + '/' + toString(date.getMonth() + 1, 2) + '/' + toString(date.getFullYear(), 4);
    return dateTimeNow;
};

interface Options {
    label: string;
    value: string;
}
interface Props {
    categories: CategoryModel[]
}
function AddTransactionForm({ categories }: Props) {
    const [form] = Form.useForm();

    return (
        <Form
            form={form}
            layout="vertical"
        >
            <>
                <Form.Item label="Categoria">
                    <Select
                        showSearch
                        placeholder='Nome da categoria'
                        optionFilterProp='children'
                        filterOption={(input, option) => (option?.label?.toString() ?? '').includes(input)}
                        filterSort={(optionA, optionB) =>
                            (optionA?.label?.toString() ?? '').toLowerCase().localeCompare((optionB?.label?.toString() ?? '').toLowerCase())
                        }
                        options={categories?.map(cat => {
                            return {
                                label: cat.name,
                                value: cat.name,
                            } as Options
                        })}
                    />
                </Form.Item>
                <Form.Item label="Descrição">
                    <Input placeholder='Gasto com o que?' />
                </Form.Item>
                <Form.Item label="Tipo">
                    <Radio.Group>
                        <Radio value="CREDIT"> Cartão de Crédito </Radio>
                        <Radio value="DEBIT"> Cartão de Dédito </Radio>
                        <Radio value="PIX"> Pix </Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item label="Valor">
                    <InputNumber placeholder='Valor da transação' addonAfter="R$" />
                </Form.Item>
            </>
            <Form.Item label="Data">
                <DatePicker defaultValue={dayjs(getTodayDate(), dateFormatList[0])} format={dateFormatList} />
            </Form.Item>

            {/* <Form.Item>
                <Button type="primary">Submit</Button>
            </Form.Item> */}
        </Form>
    )
}

export default AddTransactionForm;