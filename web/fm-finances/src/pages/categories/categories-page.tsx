import { Button, Divider, message, Space } from "antd"
import React, { useEffect, useState } from 'react';
import { PageLayout } from "src/components/page-layout"
import { PlusOutlined, UploadOutlined, RetweetOutlined } from '@ant-design/icons';
import CategoryForm from "./category-form";
import { CategoryModel } from "src/models/category";
import { useMutation, useQuery } from "react-query";
import { CategoryApi } from "src/services/category.service";
import CategoryList from "./category-list";

const CategoriesPage: React.FC = () => {
    const [isCategoryModalOpen, setIsCategoryModalOpen] = useState(false);
    const [categories, setCategories] = useState<CategoryModel[]>([])
    const { data, isError, isLoading, refetch } = useQuery("category-list", CategoryApi.getCategories);
    const [messageApi, contextHolder] = message.useMessage();

    useEffect(() => {
        if (isError) {
            messageApi.open({
                type: 'error',
                content: 'Erro ao buscar categorias',
            });
        }
    }, [isError, messageApi])


    const { mutate } = useMutation(
        () => CategoryApi.updateCategoryOrder(categories),
        {
            onSuccess: () => {
                messageApi.open({
                    type: 'success',
                    content: 'Lista de categorias ordenada com sucesso,',
                });
            },
            onError: (error) => {
                messageApi.open({
                    type: 'error',
                    content: 'Erro ao atualizar ordem de categorias',
                });
            }
        }
    )

    const showCategoryModal = () => {
        setIsCategoryModalOpen(true);
    };

    const onSuccess = () => {
        setIsCategoryModalOpen(false);
        refetch()
    };

    const onClose = () => {
        setIsCategoryModalOpen(false);
    };

    const onChange = (categories: CategoryModel[]) => {
        setCategories(categories)
    }

    const refreshList = () => {
        refetch()
    }
    return (
        <PageLayout>
            <>
                {contextHolder}
                <div className="content-layout">
                    <h1 id="page-title" className="content__title">
                        Categorias
                    </h1>
                    <div className="content__body">
                        {contextHolder}
                        <Space>
                            <Button onClick={showCategoryModal}>
                                <PlusOutlined /> Adicionar categoria
                            </Button>
                            <Button onClick={() => mutate()}>
                                <UploadOutlined /> Atualizar ordem
                            </Button>
                            <Button onClick={refreshList}>
                                <RetweetOutlined /> Atualizar lista
                            </Button>
                        </Space>
                        <Divider />

                        <CategoryList data={data} isLoading={isLoading} onChange={onChange} />

                        {isCategoryModalOpen && (
                            <CategoryForm show={isCategoryModalOpen} onClose={onClose} onSuccess={onSuccess} />
                        )}
                    </div>
                </div>
            </>
        </PageLayout>
    )
}

export default CategoriesPage;