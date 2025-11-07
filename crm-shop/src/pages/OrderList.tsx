import React from 'react';
import { Breadcrumb, Card, Form, Input, Select, Button, Tabs, Table, Tag, Empty, Dropdown } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import { Link } from 'react-router-dom';

type OrderItem = {
  id: string;
  status: 'unpaid' | 'pending' | 'shipped' | 'finished' | 'refunded' | 'deleted';
  user: { name: string; id: number };
  goods: { title: string; price: number; cover?: string };
  payMethod: string;
  payTime?: string;
  amount: number;
};

const statusMap: Record<OrderItem['status'], { label: string; color: string }> = {
  unpaid: { label: '未支付', color: 'default' },
  pending: { label: '待发货', color: 'processing' },
  shipped: { label: '待收货', color: 'warning' },
  finished: { label: '已完成', color: 'success' },
  refunded: { label: '已退款', color: 'error' },
  deleted: { label: '已删除', color: 'default' },
};

const allMock: OrderItem[] = Array.from({ length: 21 }).map((_, i) => {
  const statuses: OrderItem['status'][] = ['unpaid', 'pending', 'shipped', 'finished'];
  const st = statuses[i % statuses.length];
  return {
    id: `cp${Date.now()}${i}`,
    status: st,
    user: { name: ['张三', '李四', '王五', '赵六'][i % 4], id: 50000 + i },
    goods: { title: ['阿迪达斯鞋', 'NIKE卫衣', '华为手机壳', '小米蓝牙耳机'][i % 4], price: 99 + i },
    payMethod: st === 'unpaid' ? '--' : ['微信支付', '支付宝', '云闪付'][i % 3],
    payTime: st === 'unpaid' ? undefined : '2025-01-01 10:00:00',
    amount: 9.9 + i,
  };
});

const OrderList: React.FC = () => {
  const [form] = Form.useForm();
  const [activeTab, setActiveTab] = React.useState<string>('all');
  const [page, setPage] = React.useState<number>(1);
  const [pageSize, setPageSize] = React.useState<number>(10);

  const filtered = React.useMemo(() => {
    if (activeTab === 'all') return allMock;
    const map: Record<string, OrderItem['status']> = {
      unpaid: 'unpaid', pending: 'pending', shipped: 'shipped', finished: 'finished', refunded: 'refunded', deleted: 'deleted'
    };
    return allMock.filter(o => o.status === map[activeTab]);
  }, [activeTab]);

  const paged = React.useMemo(() => {
    const start = (page - 1) * pageSize;
    return filtered.slice(start, start + pageSize);
  }, [filtered, page, pageSize]);

  const columns: ColumnsType<OrderItem> = [
    { title: '订单编号', dataIndex: 'id', width: 220 },
    { title: '商品信息', dataIndex: 'goods', render: (g: OrderItem['goods']) => `${g.title}` },
    { title: '用户信息', dataIndex: 'user', render: (u: OrderItem['user']) => `${u.name}｜${u.id}` },
    { title: '实付金额', dataIndex: 'amount', render: (v: number) => `￥${v.toFixed(2)}` },
    { title: '支付方式', dataIndex: 'payMethod' },
    { title: '支付时间', dataIndex: 'payTime', render: (v?: string) => v || '--' },
    { title: '订单状态', dataIndex: 'status', render: (s: OrderItem['status']) => <Tag color={statusMap[s].color}>{statusMap[s].label}</Tag> },
    {
      title: '操作',
      key: 'action',
      fixed: 'right',
      width: 120,
      render: () => (
        <Dropdown
          menu={{
            items: [
              { key: 'detail', label: '订单详情' },
              { key: 'refresh', label: '刷新订单' },
            ],
            onClick: (info) => {
              if (info.key === 'refresh') {
                setPage(1);
              }
            }
          }}
        >
          <Button type="link">更多</Button>
        </Dropdown>
      )
    }
  ];

  return (
    <div>
      <Card>
        {/* 面包屑导航 */}
        <Breadcrumb style={{ marginBottom: 20 }} items={[{ title: <Link to="/home">首页</Link> }, { title: '订单管理' }, { title: '订单列表' }]} />

        {/* 顶部筛选栏：与上传图的布局风格一致 */}
        <Form form={form} layout="inline" style={{ background: '#f7f8fa', padding: 16, borderRadius: 8 }}>
          <Form.Item label="订单类型" name="type">
            <Select style={{ width: 160 }} placeholder="请选择" allowClear options={[{ value: 'all', label: '全部' }, { value: 'normal', label: '普通订单' }, { value: 'vip', label: '会员订单' }]} />
          </Form.Item>
          <Form.Item label="支付方式" name="pay">
            <Select style={{ width: 160 }} placeholder="请选择" allowClear options={[{ value: 'wechat', label: '微信支付' }, { value: 'alipay', label: '支付宝' }, { value: 'union', label: '云闪付' }]} />
          </Form.Item>
          <Form.Item label="关键词" name="kw">
            <Input style={{ width: 280 }} placeholder="商品名/订单号/用户" />
          </Form.Item>
          <Form.Item>
            <Button type="primary">查询</Button>
          </Form.Item>
        </Form>

        {/* 状态标签（使用 Tabs 实现视觉靠近上传图的顶部标签效果） */}
        <Tabs
          style={{ marginTop: 12 }}
          items={[
            { key: 'all', label: '全部' },
            { key: 'unpaid', label: '待支付' },
            { key: 'pending', label: '待发货' },
            { key: 'shipped', label: '待收货' },
            { key: 'finished', label: '已完成' },
            { key: 'refunded', label: '已退款' },
            { key: 'deleted', label: '已删除' },
          ]}
          activeKey={activeTab}
          onChange={(k) => { setActiveTab(k); setPage(1); }}
        />

        <div style={{ marginTop: 16 }} className="upload-like-box">
          <Table
            columns={columns}
            dataSource={paged}
            rowKey="id"
            pagination={{
              current: page,
              pageSize,
              total: filtered.length,
              showSizeChanger: true,
              onChange: (p, ps) => { setPage(p); setPageSize(ps); },
              pageSizeOptions: [10, 20, 50],
            }}
            locale={{ emptyText: <Empty description="暂无数据" /> }}
          />
        </div>
      </Card>
    </div>
  );
};

export default OrderList;