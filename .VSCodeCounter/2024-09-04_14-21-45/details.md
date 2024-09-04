# Details

Date : 2024-09-04 14:21:45

Directory /home/user/go/src/marketplace

Total : 78 files,  7280 codes, 759 comments, 733 blanks, all 8772 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [.env](/.env) | Properties | 5 | 0 | 0 | 5 |
| [.golangci.pipeline.yaml](/.golangci.pipeline.yaml) | YAML | 46 | 2 | 5 | 53 |
| [CHANGELOG.md](/CHANGELOG.md) | Markdown | 61 | 0 | 4 | 65 |
| [Dockerfile](/Dockerfile) | Docker | 18 | 14 | 16 | 48 |
| [Makefile](/Makefile) | Makefile | 18 | 3 | 10 | 31 |
| [README.md](/README.md) | Markdown | 57 | 0 | 3 | 60 |
| [cmd/myapp/main.go](/cmd/myapp/main.go) | Go | 11 | 8 | 4 | 23 |
| [configs/config.yaml](/configs/config.yaml) | YAML | 8 | 0 | 0 | 8 |
| [db/master/create_repluser.sql](/db/master/create_repluser.sql) | SQL | 1 | 0 | 1 | 2 |
| [db/master/master.sh](/db/master/master.sh) | Shell Script | 8 | 4 | 3 | 15 |
| [db/migrations/20240817103536_create_main_tables.sql](/db/migrations/20240817103536_create_main_tables.sql) | SQL | 48 | 6 | 1 | 55 |
| [db/migrations/20240817131113_fill_order_states.sql](/db/migrations/20240817131113_fill_order_states.sql) | SQL | 7 | 6 | 1 | 14 |
| [db/migrations/20240817132009_fill_providers.sql](/db/migrations/20240817132009_fill_providers.sql) | SQL | 7 | 6 | 1 | 14 |
| [db/migrations/20240817133433_fill_products.sql](/db/migrations/20240817133433_fill_products.sql) | SQL | 18 | 6 | 1 | 25 |
| [db/replica/replica.sh](/db/replica/replica.sh) | Shell Script | 4 | 2 | 2 | 8 |
| [docker-compose.yml](/docker-compose.yml) | YAML | 33 | 0 | 3 | 36 |
| [docs/docs.go](/docs/docs.go) | Go | 1,100 | 2 | 5 | 1,107 |
| [docs/swagger.json](/docs/swagger.json) | JSON | 1,080 | 0 | 0 | 1,080 |
| [docs/swagger.yaml](/docs/swagger.yaml) | YAML | 712 | 0 | 1 | 713 |
| [internal/app/domain/errors.go](/internal/app/domain/errors.go) | Go | 15 | 0 | 3 | 18 |
| [internal/app/domain/item.go](/internal/app/domain/item.go) | Go | 39 | 0 | 5 | 44 |
| [internal/app/domain/order.go](/internal/app/domain/order.go) | Go | 42 | 17 | 9 | 68 |
| [internal/app/domain/order_state.go](/internal/app/domain/order_state.go) | Go | 21 | 2 | 5 | 28 |
| [internal/app/domain/product.go](/internal/app/domain/product.go) | Go | 39 | 3 | 5 | 47 |
| [internal/app/domain/provider.go](/internal/app/domain/provider.go) | Go | 27 | 6 | 7 | 40 |
| [internal/app/domain/user.go](/internal/app/domain/user.go) | Go | 52 | 51 | 12 | 115 |
| [internal/app/repository/inmemrepo/imrepo_item.go](/internal/app/repository/inmemrepo/imrepo_item.go) | Go | 107 | 17 | 16 | 140 |
| [internal/app/repository/inmemrepo/imrepo_order.go](/internal/app/repository/inmemrepo/imrepo_order.go) | Go | 133 | 21 | 20 | 174 |
| [internal/app/repository/inmemrepo/imrepo_order_state.go](/internal/app/repository/inmemrepo/imrepo_order_state.go) | Go | 94 | 14 | 13 | 121 |
| [internal/app/repository/inmemrepo/imrepo_product.go](/internal/app/repository/inmemrepo/imrepo_product.go) | Go | 97 | 15 | 12 | 124 |
| [internal/app/repository/inmemrepo/imrepo_provider.go](/internal/app/repository/inmemrepo/imrepo_provider.go) | Go | 94 | 14 | 12 | 120 |
| [internal/app/repository/inmemrepo/imrepo_user.go](/internal/app/repository/inmemrepo/imrepo_user.go) | Go | 121 | 15 | 15 | 151 |
| [internal/app/repository/inmemrepo/imrepo_utils.go](/internal/app/repository/inmemrepo/imrepo_utils.go) | Go | 107 | 0 | 13 | 120 |
| [internal/app/repository/inmemrepo/inmemrepo.go](/internal/app/repository/inmemrepo/inmemrepo.go) | Go | 84 | 0 | 4 | 88 |
| [internal/app/repository/models/item.go](/internal/app/repository/models/item.go) | Go | 11 | 1 | 3 | 15 |
| [internal/app/repository/models/order.go](/internal/app/repository/models/order.go) | Go | 10 | 1 | 3 | 14 |
| [internal/app/repository/models/order_state.go](/internal/app/repository/models/order_state.go) | Go | 10 | 1 | 3 | 14 |
| [internal/app/repository/models/product.go](/internal/app/repository/models/product.go) | Go | 13 | 1 | 3 | 17 |
| [internal/app/repository/models/provider.go](/internal/app/repository/models/provider.go) | Go | 11 | 1 | 3 | 15 |
| [internal/app/repository/models/user.go](/internal/app/repository/models/user.go) | Go | 13 | 1 | 3 | 17 |
| [internal/app/repository/pgrepo/pgrepo.go](/internal/app/repository/pgrepo/pgrepo.go) | Go | 16 | 2 | 4 | 22 |
| [internal/app/repository/pgrepo/pgrepo_item.go](/internal/app/repository/pgrepo/pgrepo_item.go) | Go | 172 | 21 | 23 | 216 |
| [internal/app/repository/pgrepo/pgrepo_order.go](/internal/app/repository/pgrepo/pgrepo_order.go) | Go | 219 | 28 | 25 | 272 |
| [internal/app/repository/pgrepo/pgrepo_order_state.go](/internal/app/repository/pgrepo/pgrepo_order_state.go) | Go | 156 | 23 | 23 | 202 |
| [internal/app/repository/pgrepo/pgrepo_product.go](/internal/app/repository/pgrepo/pgrepo_product.go) | Go | 180 | 23 | 23 | 226 |
| [internal/app/repository/pgrepo/pgrepo_provider.go](/internal/app/repository/pgrepo/pgrepo_provider.go) | Go | 157 | 23 | 23 | 203 |
| [internal/app/repository/pgrepo/pgrepo_user.go](/internal/app/repository/pgrepo/pgrepo_user.go) | Go | 189 | 26 | 27 | 242 |
| [internal/app/repository/pgrepo/pgrepo_utils.go](/internal/app/repository/pgrepo/pgrepo_utils.go) | Go | 103 | 0 | 13 | 116 |
| [internal/app/services/service_interfaces.go](/internal/app/services/service_interfaces.go) | Go | 49 | 6 | 9 | 64 |
| [internal/app/services/service_item.go](/internal/app/services/service_item.go) | Go | 28 | 3 | 10 | 41 |
| [internal/app/services/service_order.go](/internal/app/services/service_order.go) | Go | 43 | 2 | 12 | 57 |
| [internal/app/services/service_order_state.go](/internal/app/services/service_order_state.go) | Go | 28 | 2 | 10 | 40 |
| [internal/app/services/service_product.go](/internal/app/services/service_product.go) | Go | 28 | 2 | 10 | 40 |
| [internal/app/services/service_provider.go](/internal/app/services/service_provider.go) | Go | 28 | 3 | 10 | 41 |
| [internal/app/services/service_token.go](/internal/app/services/service_token.go) | Go | 68 | 4 | 14 | 86 |
| [internal/app/services/service_user.go](/internal/app/services/service_user.go) | Go | 45 | 2 | 11 | 58 |
| [internal/app/transport/httpserver/auth_mw.go](/internal/app/transport/httpserver/auth_mw.go) | Go | 49 | 0 | 6 | 55 |
| [internal/app/transport/httpserver/handler_auth.go](/internal/app/transport/httpserver/handler_auth.go) | Go | 66 | 22 | 14 | 102 |
| [internal/app/transport/httpserver/handler_item.go](/internal/app/transport/httpserver/handler_item.go) | Go | 149 | 52 | 17 | 218 |
| [internal/app/transport/httpserver/handler_order.go](/internal/app/transport/httpserver/handler_order.go) | Go | 110 | 37 | 19 | 166 |
| [internal/app/transport/httpserver/handler_order_state.go](/internal/app/transport/httpserver/handler_order_state.go) | Go | 81 | 34 | 18 | 133 |
| [internal/app/transport/httpserver/handler_product.go](/internal/app/transport/httpserver/handler_product.go) | Go | 81 | 32 | 18 | 131 |
| [internal/app/transport/httpserver/handler_provider.go](/internal/app/transport/httpserver/handler_provider.go) | Go | 81 | 32 | 19 | 132 |
| [internal/app/transport/httpserver/handler_user.go](/internal/app/transport/httpserver/handler_user.go) | Go | 108 | 28 | 14 | 150 |
| [internal/app/transport/httpserver/interfaces.go](/internal/app/transport/httpserver/interfaces.go) | Go | 53 | 8 | 11 | 72 |
| [internal/app/transport/httpserver/model.go](/internal/app/transport/httpserver/model.go) | Go | 132 | 11 | 22 | 165 |
| [internal/app/transport/httpserver/password_helpers.go](/internal/app/transport/httpserver/password_helpers.go) | Go | 17 | 2 | 6 | 25 |
| [internal/app/transport/httpserver/server.go](/internal/app/transport/httpserver/server.go) | Go | 146 | 33 | 25 | 204 |
| [internal/app/transport/httpserver/services.go](/internal/app/transport/httpserver/services.go) | Go | 29 | 2 | 3 | 34 |
| [internal/app/transport/httpserver/utils.go](/internal/app/transport/httpserver/utils.go) | Go | 102 | 7 | 16 | 125 |
| [internal/app/transport/middlewares/auth.go](/internal/app/transport/middlewares/auth.go) | Go | 47 | 2 | 10 | 59 |
| [internal/app/transport/middlewares/logger.go](/internal/app/transport/middlewares/logger.go) | Go | 20 | 1 | 5 | 26 |
| [internal/app/transport/middlewares/request_id.go](/internal/app/transport/middlewares/request_id.go) | Go | 12 | 1 | 3 | 16 |
| [internal/pkg/client/client.go](/internal/pkg/client/client.go) | Go | 1 | 0 | 1 | 2 |
| [internal/pkg/config/config.go](/internal/pkg/config/config.go) | Go | 35 | 10 | 9 | 54 |
| [internal/pkg/pg/pg.go](/internal/pkg/pg/pg.go) | Go | 40 | 33 | 11 | 84 |
| [internal/pkg/utils/jwt.go](/internal/pkg/utils/jwt.go) | Go | 22 | 2 | 6 | 30 |
| [internal/pkg/utils/lvl.go](/internal/pkg/utils/lvl.go) | Go | 8 | 0 | 3 | 11 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)