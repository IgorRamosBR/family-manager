import React, { useCallback, useEffect, useRef, useState } from 'react';
import { Table } from 'antd';
import type { ColumnsType, TableProps } from 'antd/es/table';
import update from 'immutability-helper';
import { DndProvider, useDrag, useDrop } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import { CategoryModel } from '../../../models/category';

interface Props {
    data: CategoryModel[] | undefined;
    isLoading: boolean;
    onChange: (categories: CategoryModel[]) => void;
}

interface DataType {
    key: string;
    name: string;
    priority: number
    type: string;
}

interface DraggableBodyRowProps extends React.HTMLAttributes<HTMLTableRowElement> {
    index: number;
    moveRow: (dragIndex: number, hoverIndex: number) => void;
}

const type = 'DraggableBodyRow';

const DraggableBodyRow = ({
    index,
    moveRow,
    className,
    style,
    ...restProps
}: DraggableBodyRowProps) => {
    const ref = useRef<HTMLTableRowElement>(null);
    const [{ isOver, dropClassName }, drop] = useDrop({
        accept: type,
        collect: (monitor) => {
            const { index: dragIndex } = monitor.getItem() || {};
            if (dragIndex === index) {
                return {};
            }
            return {
                isOver: monitor.isOver(),
                dropClassName: dragIndex < index ? ' drop-over-downward' : ' drop-over-upward',
            };
        },
        drop: (item: { index: number }) => {
            moveRow(item.index, index);
        },
    });
    const [, drag] = useDrag({
        type,
        item: { index },
        collect: (monitor) => ({
            isDragging: monitor.isDragging(),
        }),
    });
    drop(drag(ref));

    return (
        <tr
            ref={ref}
            className={`${className}${isOver ? dropClassName : ''}`}
            style={{ cursor: 'move', ...style }}
            {...restProps}
        />
    );
};

const columns: ColumnsType<DataType> = [
    {
        title: 'Prioridade',
        dataIndex: 'priority',
        key: 'priority',
    },
    {
        title: 'Nome',
        dataIndex: 'name',
        key: 'name',
    },
    {
        title: 'Tipo',
        dataIndex: 'type',
        key: 'type',
    }
];

const mapToDataType = (categories: CategoryModel[] | undefined): DataType[] => {
    let data: DataType[] = []
    categories?.sort(compare)
    categories?.forEach((column, index) => {
        data.push({
            name: column.name,
            priority: index + 1,
            type: column.type,
            key:  (index + 1).toString()
        })
    })
    return data
}

const compare = (c1: CategoryModel, c2: CategoryModel) => {
    if ( c1.priority < c2.priority ){
        return -1;
      }
      if ( c1.priority > c2.priority ){
        return 1;
      }
      return 0;
}

const mapToCategoryModel = (columns: DataType[]): CategoryModel[] => {
    let categories: CategoryModel[] = []
    columns?.forEach((column, index) => {
        categories.push({
            name: column.name,
            type: column.type,
            priority: index + 1,
        })
    })
    return categories
} 

export function CategoryList({data, isLoading, onChange}: Props) {
    const [categories, setCategories] = useState<DataType[]>(mapToDataType(data))

    useEffect(() => {
        setCategories(mapToDataType(data))
    }, [data])

    const tableProps: TableProps<DataType> = {
        loading: isLoading,
        pagination: false,
    };

    const components = {
        body: {
            row: DraggableBodyRow,
        },
    };

    const moveRow = useCallback(
        (dragIndex: number, hoverIndex: number) => {
            const dragRow = categories[dragIndex];
            const newCategoryOrder = update(categories, {
                $splice: [
                    [dragIndex, 1],
                    [hoverIndex, 0, dragRow],
                ],
            })
            setCategories(newCategoryOrder);
            onChange(mapToCategoryModel(newCategoryOrder))
        },
        [categories, onChange],
    );

    return (
        <DndProvider backend={HTML5Backend}>
            <Table
                {...tableProps}
                columns={columns}
                dataSource={categories}
                components={components}
                onRow={(_, index) => {
                    const attr = {
                        index,
                        moveRow,
                    };
                    return attr as React.HTMLAttributes<any>;
                }}
            />
        </DndProvider>
    );
}