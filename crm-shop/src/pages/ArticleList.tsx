import React, { useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Breadcrumb, Tag } from 'antd';
import { Link } from 'react-router-dom';

type Article = { id: number; title: string; category: string; status: 'published' | 'draft' };

const initialData: Article[] = [
  { id: 1001, title: '新品发布会回顾', category: '品牌资讯', status: 'published' },
  { id: 1002, title: '校友故事：创业之路', category: '院校介绍', status: 'draft' },
];

const ArticleList: React.FC = () => {
  const [category, setCategory] = useState<string | undefined>();
  const [status, setStatus] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');

  const filtered = initialData.filter(item => {
    const byCat = category ? item.category === category : true;
    const byStatus = status ? item.status === status : true;
    const byKw = keyword ? (item.title.includes(keyword)) : true;
    return byCat && byStatus && byKw;
  });

  const columns = [
    { title: 'ID', dataIndex: 'id', width: 80 },
    { title: '文章标题', dataIndex: 'title' },
    { title: '分类', dataIndex: 'category', width: 160 },
    { title: '状态', dataIndex: 'status', width: 120, render: (s: string) => <Tag color={s === 'published' ? 'blue' : 'default'}>{s === 'published' ? '已发布' : '草稿'}</Tag> },
    { title: '操作', dataIndex: 'action', width: 200, render: () => <div style={{ display: 'flex', gap: 8 }}><Button type="link">编辑</Button><Button type="link" danger>删除</Button></div> }
  ];

  return (
    <div>
      <Card>
        <Breadcrumb style={{ marginBottom: 20 }}>
          <Breadcrumb.Item>
            <Link to="/home">首页</Link>
          </Breadcrumb.Item>
          <Breadcrumb.Item>内容管理</Breadcrumb.Item>
          <Breadcrumb.Item>文章列表</Breadcrumb.Item>
        </Breadcrumb>

        <Form layout="inline" style={{ background: '#f7f8fa', padding: 16, borderRadius: 8 }}>
          <Form.Item label="文章分类">
            <Select
              style={{ width: 220 }}
              placeholder="请选择"
              value={category}
              onChange={setCategory}
              options={[{ value: '品牌资讯', label: '品牌资讯' }, { value: '院校介绍', label: '院校介绍' }]}
              allowClear
            />
          </Form.Item>
          <Form.Item label="发布状态">
            <Select
              style={{ width: 180 }}
              placeholder="请选择"
              value={status}
              onChange={setStatus}
              options={[{ value: 'published', label: '已发布' }, { value: 'draft', label: '草稿' }]}
              allowClear
            />
          </Form.Item>
          <Form.Item label="关键词">
            <Input
              style={{ width: 280 }}
              placeholder="请输入关键词"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
            />
          </Form.Item>
          <Form.Item>
            <Button type="primary">查询</Button>
          </Form.Item>
        </Form>

        <div style={{ marginTop: 12, display: 'flex', justifyContent: 'flex-start' }}>
          <Button type="primary" size="small">新增文章</Button>
        </div>

        <div style={{ marginTop: 16 }}>
          <Table
            columns={columns}
            dataSource={filtered}
            pagination={false}
            locale={{ emptyText: <Empty description="暂无数据" /> }}
            rowKey="id"
          />
        </div>
      </Card>
    </div>
  );
};

export default ArticleList;