import HomePage from "pages/homepage";

import Icon from "@mui/material/Icon";

const routes = [
    {
        type: "collapse",
        name: "homepage",
        key: "homepage",
        icon: <Icon fontSize="small">homepage</Icon>,
        route: "/dashboard",
        component: <HomePage />
    }
]


export default routes
