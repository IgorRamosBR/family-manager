import { Button, Divider, message, Space } from 'antd'
import { PlusOutlined, UploadOutlined, RetweetOutlined } from '@ant-design/icons';
import React, { useEffect, useState } from 'react'
import AddCategoryModal from './add-category/add-category-modal';
import { useMutation, useQuery } from 'react-query';
import { CategoryApi } from '../../api/category';
import { CategoryList } from './list-categories/list-categories';
import { CategoryModel } from '../../models/category';

const Category: React.FC = () => {
    const [isCategoryModalOpen, setIsCategoryModalOpen] = useState(false);
    const [categories, setCategories] = useState<CategoryModel[]>([])
    const { data, isError, isLoading, refetch } = useQuery("category-list", CategoryApi.getCategories);
    const [messageApi, contextHolder] = message.useMessage();

    //const queryClient = useQueryClient();

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
            onError:(error) => {
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
        <div>
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

            <CategoryList data={data} isLoading={isLoading} onChange={onChange}/>

            {isCategoryModalOpen && (
                <AddCategoryModal show={isCategoryModalOpen} onClose={onClose} onSuccess={onSuccess}/>
            )}
        </div>
    )

}

export default Category;