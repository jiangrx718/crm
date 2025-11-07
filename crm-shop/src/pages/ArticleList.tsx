import React, { useMemo, useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Breadcrumb, Image, Popconfirm, message } from 'antd';
import { Link } from 'react-router-dom';

type Article = {
  id: number;
  title: string;
  category: string;
  cover: string;
  relatedGoods: string;
  views: number;
  time: string; // YYYY-MM-DD HH:mm
  status: 'published' | 'draft';
};

const categories = ['品牌资讯', '生活家居', '潮流文化', '🎧分类'];

const initialData: Article[] = Array.from({ length: 36 }, (_, i) => {
  const id = 237 + i;
  const cat = categories[i % categories.length];
  const titlePool = [
    '电影评谈 “618” 回归｜破圈新风尚',
    '联博观察｜考究美学迈向文化潮新时代',
    '鉴宇｜国内外KOL，初创团队评审会吵',
    '把温柔的日子放在盘里',
    '街头艺术周刊｜跨界装置展精选',
    '球鞋文化速递｜热门联名一览',
  ];
  return {
    id,
    title: titlePool[i % titlePool.length],
    category: cat,
    cover: `https://picsum.photos/seed/a${id}/60/60`,
    relatedGoods: i % 2 === 0 ? 'Kaleidos 万花筒装饰画合集' : '联名限量周边',
    views: 200 + (i * 7) % 1300,
    time: `2025-04-${String(1 + (i % 9)).padStart(2, '0')} 16:${String(20 + (i % 40)).padStart(2, '0')}`,
    status: i % 3 === 0 ? 'draft' : 'published',
  };
});

const ArticleList: React.FC = () => {
  const [category, setCategory] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');
  const [data, setData] = useState<Article[]>(initialData);
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);

  const filtered = useMemo(() => (
    data.filter(item => {
      const byCat = category ? item.category === category : true;
      const byKw = keyword ? item.title.includes(keyword) : true;
      return byCat && byKw;
    })
  ), [data, category, keyword]);

  const paged = useMemo(() => (
    filtered.slice((page - 1) * pageSize, page * pageSize)
  ), [filtered, page, pageSize]);

  const togglePublish = (id: number) => {
    setData(prev => prev.map(a => a.id === id ? { ...a, status: a.status === 'published' ? 'draft' : 'published' } : a));
  };

  const removeById = (id: number) => {
    setData(prev => prev.filter(a => a.id !== id));
    message.success('已删除文章');
  };

  const copyLink = async (id: number) => {
    const url = `https://example.com/article/${id}`;
    try {
      await navigator.clipboard.writeText(url);
      message.success('链接已复制');
    } catch {
      message.info(url);
    }
  };

  const columns = [
    { title: 'ID', dataIndex: 'id', width: 80 },
    { title: '文章图片', dataIndex: 'cover', width: 100, render: (src: string) => <Image src={src} width={40} height={40} /> },
    { title: '文章名称', dataIndex: 'title' },
    { title: '关联商品', dataIndex: 'relatedGoods' },
    { title: '浏览量', dataIndex: 'views', width: 100 },
    { title: '时间', dataIndex: 'time', width: 180 },
    { title: '操作', dataIndex: 'action', width: 260, render: (_: any, record: Article) => (
      <div style={{ display: 'flex', gap: 8 }}>
        <Button type="link">编辑</Button>
        <Button type="link" onClick={() => togglePublish(record.id)}>
          {record.status === 'published' ? '取消发布' : '发布'}
        </Button>
        <Popconfirm
          title="确认删除当前文章吗？"
          okText="删除"
          cancelText="取消"
          okButtonProps={{ danger: true }}
          onConfirm={() => removeById(record.id)}
        >
          <Button type="link" danger>删除</Button>
        </Popconfirm>
        <Button type="link" onClick={() => copyLink(record.id)}>复制链接</Button>
      </div>
    ) }
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
              options={categories.map(c => ({ value: c, label: c }))}
              allowClear
            />
          </Form.Item>
          <Form.Item label="文章搜索">
            <Input
              style={{ width: 280 }}
              placeholder="请输入"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
            />
          </Form.Item>
          <Form.Item>
            <Button type="primary">查询</Button>
          </Form.Item>
        </Form>

        <div style={{ marginTop: 12, display: 'flex', justifyContent: 'flex-start' }}>
          <Button type="primary" size="small">添加文章</Button>
        </div>

        <div style={{ marginTop: 16 }}>
          <Table
            columns={columns}
            dataSource={paged}
            pagination={{
              current: page,
              pageSize,
              total: filtered.length,
              showSizeChanger: true,
              pageSizeOptions: [10, 20, 50],
              showTotal: (total) => `共 ${total} 条`,
              onChange: (p, ps) => { setPage(p); setPageSize(ps); },
            }}
            locale={{ emptyText: <Empty description="暂无数据" /> }}
            rowKey="id"
          />
        </div>
      </Card>
    </div>
  );
};

export default ArticleList;