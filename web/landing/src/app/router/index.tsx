// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { LazyExoticComponent } from 'react';
import { Switch } from 'react-router-dom';

const WelcomePage = React.lazy(() => import('@components/WelcomePage'));
const SignIn = React.lazy(() => import('@/app/views/SignIn'));
const SignUp = React.lazy(() => import('@/app/views/SignUp'));
const ChangePassword = React.lazy(() => import('@/app/views/ChangePassword'));
const ConfirmEmail = React.lazy(() => import('@/app/views/ConfirmEmail'));
const RecoverPassword = React.lazy(() => import('@/app/views/RecoverPassword'));

export interface RouteItem {
    path: string,
    component: React.FC<any>,
    exact: boolean,
    children?: ComponentRoutes[],
    with?: (
        child: ComponentRoutes,
        parrent: ComponentRoutes
    ) => ComponentRoutes,
    addChildren?: (children: ComponentRoutes[]) => ComponentRoutes
}

/** Route base config implementation */
export class ComponentRoutes implements RouteItem {
    /** data route config*/
    constructor(
        public path: string,
        public component: React.FC |
            LazyExoticComponent<React.FC<{ children: ComponentRoutes[] }>>,
        public exact: boolean,
        public children?: ComponentRoutes[],
    ) { };
    /* change path for children routes */
    public with(
        child: ComponentRoutes,
        parrent: ComponentRoutes
    ): ComponentRoutes {
        child.path = `${parrent.path}/${child.path}`;

        return this;
    };
    /* adds children routes to route */
    public addChildren(children: ComponentRoutes[]): ComponentRoutes {
        this.children = children.map((child) => child.with(child, this));

        return this;
    };
};

/** Route config implementation */
export class RouteConfig {
    public static WelcomePage: ComponentRoutes = new ComponentRoutes(
        '/',
        WelcomePage,
        true,
    );
    public static SignIn: ComponentRoutes = new ComponentRoutes(
        '/sign-in',
        SignIn,
        true
    );
    public static SignUp: ComponentRoutes = new ComponentRoutes(
        '/sign-up',
        SignUp,
        true
    );
    public static ResetPassword: ComponentRoutes = new ComponentRoutes(
        '/change-password',
        ChangePassword,
        true
    );
    public static ConfirmEmail: ComponentRoutes = new ComponentRoutes(
        '/email/confirm',
        ConfirmEmail,
        true,
    );
    public static RecoverPassword: ComponentRoutes = new ComponentRoutes(
        '/recover-password',
        RecoverPassword,
        true,
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.WelcomePage,
        RouteConfig.SignIn,
        RouteConfig.SignUp,
        RouteConfig.ResetPassword,
        RouteConfig.ConfirmEmail,
        RouteConfig.RecoverPassword,
    ];
};

export const Route: React.FC<RouteItem> = ({
    component: Component, ...children
}) =>
    <Component {...children} />;

export const Routes = () => {
    return (
        <Switch>
            {
                RouteConfig.routes.map((route, index) => (
                    <Route
                        key={index}
                        path={route.path}
                        component={route.component}
                        exact={route.exact}
                        children={route.children}
                    />
                )
                )
            }
        </Switch>
    );
};
