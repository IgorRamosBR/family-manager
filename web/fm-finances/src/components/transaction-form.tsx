import { Button, DatePicker, Form, FormInstance, Input, InputNumber, message, Radio, Select } from 'antd';
import React, { useState } from 'react';
import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import { CategoryModel } from 'src/models/category';
import { useMutation } from 'react-query';
import { TransactionApi } from 'src/services/transaction.service';
import { TransactionModel } from 'src/models/transaction';

dayjs.extend(customParseFormat);
const dateFormatList = ['DD/MM/YYYY', 'DD/MM/YY'];

interface Options {
    label: string;
    value: string;
}
interface Props {
    categories: CategoryModel[]
    onFinish: () => void
}
const TransactionForm: React.FC<Props> = ({ categories, onFinish }) => {
    const formRef = React.useRef<FormInstance>(null);
    const [transaction, setTransaction] = useState<TransactionModel>({} as TransactionModel)
    const [messageApi, contextHolder] = message.useMessage();

    const handleChange = (name: string, value: any) => {
     if (name === 'category') {
        value = categories?.find(cat => cat.name === value)
     }
     setTransaction((prev) => {
            return { ...prev, [name]: value }
        });
    }

    const handleSubmit = () => {
        if (!transaction.category) {
            displayErrorMessage('Categoria é obrigatória')
            return
        }
        if (!transaction.paymentMethod) {
            displayErrorMessage('Tipo da transação é obrigatório')
            return
        }
        if (!transaction.value) {
            displayErrorMessage('Valor é obrigatório')
            return
        }
        if (!transaction.date) {
            displayErrorMessage('Data é obrigatória')
            return
        }

        mutate()
    }

    const displayErrorMessage = (message: string) => {
        messageApi.open({
            type: 'error',
            content: message,
          });
    }

    const handleCancel = () => {
        formRef.current?.resetFields();
    }

    const { mutate } = useMutation(
        () => TransactionApi.createTransaction(transaction),
        {
            onSuccess: () => {
                formRef.current?.resetFields();
                onFinish()
            },
            onError: (error) => {
                console.log(error)
                messageApi.open({
                    type: 'error',
                    content: 'Erro ao salvar transação',
                });
            }
        }
    )

    return (
        <>
        {contextHolder}
        <Form
            layout="vertical"
            ref={formRef}
        >
            <>
                <Form.Item label="Categoria" rules={[{ required: true, message: 'Por favor, selecione uma categoria!' }]}>
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
                        onChange={(v) => handleChange('category', v)}
                        
                    />
                </Form.Item>
                <Form.Item label="Descrição">
                    <Input placeholder='Gasto com o que?' onChange={(v) => handleChange('description', v.target.value)}/>
                </Form.Item>
                <Form.Item label="Meio de pagamento" rules={[{ required: true, message: 'Por favor, selecione um tipo!' }]}>
                    <Radio.Group onChange={(v) => handleChange('paymentMethod', v.target.value)}>
                        <Radio value="CREDIT_CARD"> Cartão de Crédito </Radio>
                        <Radio value="DEBIT_CARD"> Cartão de Dédito </Radio>
                        <Radio value="PIX"> Pix </Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item label="Valor" rules={[{ required: true, message: 'Por favor, insira um valor!' }]}>
                    <InputNumber placeholder='Valor da transação' addonBefore="R$" step={"0.01"} onChange={(v) => handleChange('value', v)}/>
                </Form.Item>
            </>
            <Form.Item label="Data" rules={[{ required: true, message: 'Por favor, insira uma data!' }]}>
                <DatePicker format={dateFormatList} name='date' onChange={(v) => handleChange('date', v?.format(dateFormatList[0]))}/>
            </Form.Item>

            <div style={{ display: 'flex', justifyContent: 'right' }}>
                <Button type="primary" style={{ marginRight: '15px' }} onClick={handleSubmit}>Salvar</Button>
                <Button onClick={handleCancel}>Cancelar</Button>
            </div>
        </Form>
        </>
        
    )
}

export default TransactionForm;