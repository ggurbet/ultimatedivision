// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { lazy } from 'react';
import { Switch } from 'react-router-dom';

const MarketPlace = lazy(() => import('@/app/views/MarketPlacePage'));
const Club = lazy(() => import('@/app/views/ClubPage'));
const FootballerCard = lazy(() => import('@/app/views/FootballerCardPage'));
const FootballField = lazy(() => import('@/app/views/FootballFieldPage'));
const WhitePaper = lazy(() => import('@/app/views/WhitePaperPage'));
const Tokenomics = lazy(() => import('@/app/views/TokenomicsPage'));
const Store = lazy(() => import('@/app/views/StorePage'));

import Summary from '@components/WhitePaper/Summary';
import GameMechanics from '@components/WhitePaper/GameMechanics';
import PayToEarnEconomy from '@components/WhitePaper/PayToEarnEconomy';
import Technology from '@components/WhitePaper/Technology';
import Fund from '@components/Tokenomics/Fund';
import PlayToEarn from '@components/Tokenomics/PlayToEarn';
import Spending from '@components/Tokenomics/Spending';
import Staking from '@components/Tokenomics/Staking';

/** Route base config implementation */
export class ComponentRoutes {
    /** data route config*/
    constructor(
        public path: string,
        public component: React.FC<any>,
        public exact: boolean,
        public children?: ComponentRoutes[]
    ) { }
    /** Method for creating child subroutes path */
    public with(child: ComponentRoutes, parrent: ComponentRoutes): ComponentRoutes {
        child.path = `${parrent.path}/${child.path}`;

        return this;
    }
    /** Call with method for each child */
    public addChildren(children: ComponentRoutes[]): ComponentRoutes {
        this.children = children.map(item => item.with(item, this));

        return this;
    }
};

/** interfafe fot AboutPage subroutes */
interface RouteItem {
    path: string;
    component: React.FC<any>;
    exact: boolean;
    children?: ComponentRoutes[];
    with?: (child: ComponentRoutes, parrent: ComponentRoutes) => ComponentRoutes;
    addChildren?: (children: ComponentRoutes[]) => ComponentRoutes;
}

/** Route config implementation */
export class RouteConfig {
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        '/marketplace',
        MarketPlace,
        true,
    );
    public static FootballerCard: ComponentRoutes = new ComponentRoutes(
        '/card',
        FootballerCard,
        true,
    );
    public static FootballField: ComponentRoutes = new ComponentRoutes(
        '/field',
        FootballField,
        true,
    );
    public static Store: ComponentRoutes = new ComponentRoutes(
        '/store',
        Store,
        true,
    );
    public static Club: ComponentRoutes = new ComponentRoutes(
        '/club',
        Club,
        true,
    );
    public static Whitepaper: ComponentRoutes = new ComponentRoutes(
        '/whitepaper',
        WhitePaper,
        false
    );
    public static Tokenomics: ComponentRoutes = new ComponentRoutes(
        '/tokenomics',
        Tokenomics,
        false
    );
    public static Summary: ComponentRoutes = new ComponentRoutes(
        'summary',
        Summary,
        true
    );
    public static GameMechanick: ComponentRoutes = new ComponentRoutes(
        'game-mechanicks',
        GameMechanics,
        true
    );
    public static PayToEarnEconomy: ComponentRoutes = new ComponentRoutes(
        'pay-to-earn-and-economy',
        PayToEarnEconomy,
        true
    );
    public static Technology: ComponentRoutes = new ComponentRoutes(
        'technology',
        Technology,
        true
    );
    public static Spending: ComponentRoutes = new ComponentRoutes(
        'udt-spending',
        Spending,
        true
    );
    public static PayToEarn: ComponentRoutes = new ComponentRoutes(
        'pay-to-earn',
        PlayToEarn,
        true
    );
    public static Staking: ComponentRoutes = new ComponentRoutes(
        'staking',
        Staking,
        true
    );
    public static Fund: ComponentRoutes = new ComponentRoutes(
        'ud-dao-fund',
        Fund,
        true
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        '/',
        MarketPlace,
        true,
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.Default,
        RouteConfig.FootballField,
        RouteConfig.MarketPlace,
        RouteConfig.Club,
        RouteConfig.FootballerCard,
        RouteConfig.Store,
        RouteConfig.Whitepaper.addChildren([
            RouteConfig.Summary,
            RouteConfig.GameMechanick,
            RouteConfig.PayToEarnEconomy,
            RouteConfig.Technology,
        ]),
        RouteConfig.Tokenomics.addChildren([
            RouteConfig.Spending,
            RouteConfig.PayToEarn,
            RouteConfig.Staking,
            RouteConfig.Fund,
        ]),
    ];
};

export const Route: React.FC<RouteItem> = ({
    component: Component,
    ...children
}) => <Component {...children} />;

export const Routes = () =>
    <Switch>
        {RouteConfig.routes.map((route, index) =>
            <Route
                key={index}
                path={route.path}
                component={route.component}
                exact={route.exact}
                children={route.children}
            />
        )}
    </Switch>;
