/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { lazy } from 'react';
import { RouteProps, Switch } from 'react-router-dom';

const FootballerCard = lazy(() => import('../components/FootballerCardPage/FootballerCard'));
const FootballField = lazy(() => import('../components/FootballFieldPage/FootballField'));
const MarketPlace = lazy(() => import('../components/MarketPlacePage/MarketPlace'));

/** Route base config implementation */
export class ComponentRoutes {
    /** data route config*/
    constructor(
        public path: string,
        public component: React.FC,
        public exact: boolean,
    ) {
        this.path = path;
        this.component = component;
        this.exact = exact;
    }
};

/** Route config implementation */
export class RouteConfig {
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        '/ud/marketplace',
        MarketPlace,
        true,
    );
    public static FootballerCard: ComponentRoutes = new ComponentRoutes(
        '/ud/marketplace/card',
        FootballerCard,
        true,
    );
    public static FootballField: ComponentRoutes = new ComponentRoutes(
        '/ud/field',
        FootballField,
        true,
    );
    public static MyCards: ComponentRoutes = new ComponentRoutes(
        '/ud/marketplace/club',
        MarketPlace,
        true,
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        '/ud/',
        MarketPlace,
        true,
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.MarketPlace,
        RouteConfig.FootballerCard,
        RouteConfig.FootballField,
        RouteConfig.MyCards,
        RouteConfig.Default,
    ];
};

type RoutesProps = { component: React.FC } & RouteProps;

const Route: React.FC<RoutesProps> = ({
    component: Component, ...children
}) =>
    <Component {...children} />;

export const Routes = () =>
    <Switch>
        {RouteConfig.routes.map((route, index) =>
            <Route
                key={index}
                path={route.path}
                component={route.component}
                exact={route.exact}
            />,
        )}
    </Switch>;
