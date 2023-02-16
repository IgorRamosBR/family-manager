import React from 'react';
import { BankOutlined, HomeOutlined, TableOutlined, CheckOutlined } from '@ant-design/icons';
import { Layout, Menu, theme, Typography } from 'antd';
import './home.css'
import { BrowserRouter as Router, Switch, Link } from 'react-router-dom';
import Budget from '../budget/budget';
import Category from '../category/category';
import LoginButton from '../../components/login/login';
import { useAuth0 } from '@auth0/auth0-react';
import { Callback } from '../../components/callback/callback';
import { ProtectedRoute } from '../../components/auth/protected-route';
const { Text } = Typography;

const { Header, Content, Sider } = Layout;
const menu = [
    { name: <Link to="/">Início</Link>, icon: HomeOutlined },
    { name: <Link to="/budget">Controle familiar</Link>, icon: TableOutlined },
    { name: <Link to="/categories">Categorias</Link>, icon: TableOutlined },
    { name: "Investimentos", icon: BankOutlined },
    { name: "Metas", icon: CheckOutlined },
]

const routes = [
    {
        path: "/",
        exact: true,
        main: () => <Budget />
    },
    {
        path: "/budget",
        exact: true,
        main: () => <Budget />
    },
    {
        path: "/categories",
        exact: true,
        main: () => <Category />
    },
    {
        path: "/callback",
        exact: true,
        main: () => <Callback />
    }
];

const Home: React.FC = () => {
    const {
        token: { colorBgContainer },
    } = theme.useToken();
    const { user, isAuthenticated } = useAuth0();
      
    return (
        <Router>
            <Layout>
                {isAuthenticated && <Sider
                    breakpoint="lg"
                    collapsedWidth="0"
                    onBreakpoint={(broken) => {
                        console.log(broken);
                    }}
                    onCollapse={(collapsed, type) => {
                        console.log(collapsed, type);
                    }}
                >
                    <div className="logo">
                        <span style={{ color: 'white' }}>
                            Olá, {user?.name}
                        </span>

                    </div>
                    <Menu
                        theme="dark"
                        mode="inline"
                        defaultSelectedKeys={['1']}
                        items={menu.map(
                            (item, index) => ({
                                key: String(index + 1),
                                icon: React.createElement(item.icon),
                                label: item.name,
                            }),
                        )}
                    />
                </Sider>
                }
                <Layout>
                    <Header style={{ padding: 0, background: colorBgContainer }}>
                        <span style={{ paddingLeft: 40, fontSize: '24px' }}>Finance manager</span>
                    </Header>
                    <Content style={{ margin: '24px 16px 0', height: '800px' }}>
                        {!isAuthenticated ?
                            <>
                                <div style={{ display: 'flex', alignContent: 'center', justifyContent: 'center' }}>
                                    <Text strong>Faça o login antes de utilizar esse site.</Text>

                                </div>
                                <div style={{ display: 'flex', alignContent: 'center', justifyContent: 'center' }}>
                                    <LoginButton />
                                </div>
                            </> :
                            <Switch>
                                {routes.map((route, index) => (
                                    <ProtectedRoute
                                        key={index}
                                        path={route.path}
                                        exact={route.exact}
                                        component={route.main}
                                    />
                                ))}
                            </Switch>
                        }
                    </Content>
                </Layout>
            </Layout>
        </Router>
    );
};

export default Home;