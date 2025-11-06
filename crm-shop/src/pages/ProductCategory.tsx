import React, { useMemo, useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Image, Tag, Breadcrumb } from 'antd';
import { Link } from 'react-router-dom';

const ProductCategory: React.FC = () => {
  const [categoryId, setCategoryId] = useState<string | undefined>();
  const [status, setStatus] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');

  type Cat = {
    id: number;
    name: string;
    icon?: string;
    sort: number;
    status: 'enabled' | 'disabled';
    children?: Cat[];
  };

  const mockData: Cat[] = [
    {
      id: 7,
      name: '生活家居',
      icon: 'https://via.placeholder.com/32?text=家',
      sort: 999,
      status: 'enabled',
      children: [
        { id: 701, name: '收纳整理', icon: 'https://via.placeholder.com/32?text=收', sort: 30, status: 'enabled' },
        { id: 702, name: '床上用品', icon: 'https://via.placeholder.com/32?text=床', sort: 20, status: 'enabled', children: [
          { id: 70201, name: '被子', icon: 'https://via.placeholder.com/32?text=被', sort: 10, status: 'enabled' },
          { id: 70202, name: '枕头', icon: 'https://via.placeholder.com/32?text=枕', sort: 9, status: 'disabled' }
        ]},
      ]
    },
    {
      id: 53,
      name: '运动专区',
      icon: 'https://via.placeholder.com/32?text=运',
      sort: 995,
      status: 'enabled',
      children: [
        { id: 5301, name: '跑步', icon: 'https://via.placeholder.com/32?text=跑', sort: 12, status: 'enabled' },
        { id: 5302, name: '健身', icon: 'https://via.placeholder.com/32?text=健', sort: 11, status: 'enabled' },
        { id: 5303, name: '户外运动', icon: 'https://via.placeholder.com/32?text=外', sort: 10, status: 'disabled' },
        { id: 5304, name: '球类运动', icon: 'https://via.placeholder.com/32?text=球', sort: 8, status: 'enabled' }
      ]
    },
    {
      id: 8,
      name: '电子产品',
      icon: 'https://via.placeholder.com/32?text=电',
      sort: 17,
      status: 'enabled',
      children: [
        { id: 801, name: '手机', icon: 'https://via.placeholder.com/32?text=机', sort: 10, status: 'enabled' },
        { id: 802, name: '电脑', icon: 'https://via.placeholder.com/32?text=脑', sort: 9, status: 'enabled' },
        { id: 803, name: '智能穿戴', icon: 'https://via.placeholder.com/32?text=戴', sort: 8, status: 'disabled' },
      ]
    },
    {
      id: 1,
      name: '家用电器',
      icon: 'https://via.placeholder.com/32?text=器',
      sort: 9,
      status: 'enabled',
      children: [
        { id: 101, name: '冰箱', icon: 'https://via.placeholder.com/32?text=冰', sort: 10, status: 'enabled' },
        { id: 102, name: '洗衣机', icon: 'https://via.placeholder.com/32?text=洗', sort: 9, status: 'enabled' },
        { id: 103, name: '空调', icon: 'https://via.placeholder.com/32?text=调', sort: 8, status: 'disabled' },
      ]
    },
    {
      id: 3,
      name: '家具装饰',
      icon: 'https://via.placeholder.com/32?text=装',
      sort: 7,
      status: 'enabled',
      children: [
        { id: 301, name: '沙发', icon: 'https://via.placeholder.com/32?text=沙', sort: 10, status: 'enabled' },
        { id: 302, name: '餐桌', icon: 'https://via.placeholder.com/32?text=桌', sort: 9, status: 'enabled' },
        { id: 303, name: '灯具', icon: 'https://via.placeholder.com/32?text=灯', sort: 8, status: 'disabled' },
      ]
    },
    {
      id: 6,
      name: '美妆护肤',
      icon: 'https://via.placeholder.com/32?text=妆',
      sort: 6,
      status: 'enabled',
      children: [
        { id: 601, name: '面膜', icon: 'https://via.placeholder.com/32?text=膜', sort: 10, status: 'enabled' },
        { id: 602, name: '护肤乳', icon: 'https://via.placeholder.com/32?text=乳', sort: 9, status: 'enabled' },
        { id: 603, name: '口红', icon: 'https://via.placeholder.com/32?text=红', sort: 8, status: 'disabled' },
      ]
    },
    {
      id: 4,
      name: '居家餐厨',
      icon: 'https://via.placeholder.com/32?text=厨',
      sort: 6,
      status: 'enabled',
      children: [
        { id: 401, name: '锅具', icon: 'https://via.placeholder.com/32?text=锅', sort: 10, status: 'enabled' },
        { id: 402, name: '餐具', icon: 'https://via.placeholder.com/32?text=餐', sort: 9, status: 'enabled' },
        { id: 403, name: '刀具', icon: 'https://via.placeholder.com/32?text=刀', sort: 8, status: 'disabled' },
      ]
    },
    {
      id: 2,
      name: '电视影音',
      icon: 'https://via.placeholder.com/32?text=视',
      sort: 3,
      status: 'enabled',
      children: [
        { id: 201, name: '电视机', icon: 'https://via.placeholder.com/32?text=TV', sort: 10, status: 'enabled' },
        { id: 202, name: '音响', icon: 'https://via.placeholder.com/32?text=音', sort: 9, status: 'enabled' },
        { id: 203, name: '投影', icon: 'https://via.placeholder.com/32?text=影', sort: 8, status: 'disabled' },
      ]
    },
    {
      id: 9,
      name: '日用文创',
      icon: 'https://via.placeholder.com/32?text=文',
      sort: 1,
      status: 'enabled',
      children: [
        { id: 901, name: '文具', icon: 'https://via.placeholder.com/32?text=具', sort: 10, status: 'enabled' },
        { id: 902, name: '礼品', icon: 'https://via.placeholder.com/32?text=礼', sort: 9, status: 'enabled' },
        { id: 903, name: '创意摆件', icon: 'https://via.placeholder.com/32?text=件', sort: 8, status: 'disabled' },
      ]
    },
  ];

  const dataSource = useMemo(() => {
    const filterByKeyword = (items: Cat[]): Cat[] =>
      items
        .filter((item) => (status ? item.status === status : true))
        .filter((item) => (keyword ? item.name.includes(keyword) : true))
        .map((item) => ({
          ...item,
          children: item.children ? filterByKeyword(item.children) : undefined,
        }));

    const filtered = filterByKeyword(mockData);
    if (!categoryId) return filtered;
    // 简单根据顶级分类ID过滤。
    return filtered.filter((item) => String(item.id) === String(categoryId));
  }, [categoryId, status, keyword]);

  const categoryOptions = useMemo(() => mockData.map((c) => ({ value: String(c.id), label: c.name })), []);

  const columns = [
    { title: 'ID', dataIndex: 'id', width: 80 },
    { title: '分类名称', dataIndex: 'name' },
    { title: '分类图标', dataIndex: 'icon', render: (src: string) => src ? <Image src={src} width={32} height={32} /> : '-' },
    { title: '排序', dataIndex: 'sort', width: 100 },
    { title: '状态', dataIndex: 'status', width: 120, render: (s: string) => <Tag color={s === 'enabled' ? 'blue' : 'default'}>{s === 'enabled' ? '开启' : '关闭'}</Tag> },
    { title: '操作', dataIndex: 'action', width: 160, render: () => <div style={{ display: 'flex', gap: 8 }}><Button type="link">编辑</Button><Button type="link" danger>删除</Button></div> }
  ];

  return (
    <div>
      <Card>
        {/* 面包屑导航 */}
        <Breadcrumb style={{ marginBottom: 20 }}>
          <Breadcrumb.Item>
            <Link to="/home">首页</Link>
          </Breadcrumb.Item>
          <Breadcrumb.Item>商品管理</Breadcrumb.Item>
          <Breadcrumb.Item>商品分类</Breadcrumb.Item>
        </Breadcrumb>

        <Form layout="inline" style={{ background: '#f7f8fa', padding: 16, borderRadius: 8 }}>
          <Form.Item label="商品分类">
            <Select
              style={{ width: 220 }}
              placeholder="请选择"
              value={categoryId}
              onChange={setCategoryId}
              options={categoryOptions}
              allowClear
            />
          </Form.Item>
          <Form.Item label="分类状态">
            <Select
              style={{ width: 180 }}
              placeholder="请选择"
              value={status}
              onChange={setStatus}
              options={[{ value: 'enabled', label: '启用' }, { value: 'disabled', label: '禁用' }]}
              allowClear
            />
          </Form.Item>
          <Form.Item label="分类名称">
            <Input
              style={{ width: 280 }}
              placeholder="请输入分类名称"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
            />
          </Form.Item>
          <Form.Item>
            <Button type="primary">查询</Button>
          </Form.Item>
        </Form>

        <div style={{ marginTop: 12, display: 'flex', justifyContent: 'flex-start' }}>
          <Button type="primary" size="small">添加分类</Button>
        </div>

        <div style={{ marginTop: 16 }}>
          <Table
            columns={columns}
            dataSource={dataSource}
            pagination={false}
            locale={{ emptyText: <Empty description="暂无数据" /> }}
            rowKey="id"
            expandable={{
              indentSize: 20,
              rowExpandable: (record: Cat) => Array.isArray(record.children) && record.children.length > 0,
            }}
          />
        </div>
      </Card>
    </div>
  );
};

export default ProductCategory;