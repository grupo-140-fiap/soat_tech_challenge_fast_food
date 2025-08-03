# Makefile para Fast Food API

.PHONY: help build docker-build docker-run k8s-deploy k8s-delete k8s-status test clean

# Variáveis
APP_NAME=fast-food-api
DOCKER_IMAGE=$(APP_NAME):latest
NAMESPACE=fast-food

help: ## Mostra este help
	@echo "Fast Food API - Comandos disponíveis:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Compila a aplicação Go
	@echo "🔨 Compilando aplicação..."
	go mod tidy
	go build -o bin/$(APP_NAME) cmd/server/main.go

docker-build: ## Constrói a imagem Docker
	@echo "🐳 Construindo imagem Docker..."
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Executa com Docker Compose
	@echo "🚀 Iniciando aplicação com Docker Compose..."
	docker compose up --build -d
	@echo "✅ Aplicação disponível em: http://localhost:8080"
	@echo "📚 Swagger disponível em: http://localhost:8080/swagger/index.html"

docker-stop: ## Para containers Docker Compose
	@echo "🛑 Parando containers..."
	docker compose down

helm-deploy: docker-build ## Deploy com Helm
	@echo "🚀 Fazendo deploy com Helm..."
	helm upgrade --install $(APP_NAME) ./helm/fast-food --create-namespace
	@echo "⏳ Aguardando pods ficarem prontos..."
	kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=fast-food -n $(NAMESPACE) --timeout=300s
	@echo "✅ Deploy concluído!"

helm-delete: ## Remove release Helm
	@echo "🗑️ Removendo release Helm..."
	helm uninstall $(APP_NAME) --namespace $(NAMESPACE) || true
	kubectl delete namespace $(NAMESPACE) --ignore-not-found=true

helm-status: ## Verifica status do release Helm
	@echo "📊 Status do release Helm:"
	@echo ""
	@echo "🏷️ Release:"
	helm status $(APP_NAME) -n $(NAMESPACE) 2>/dev/null || echo "Release não encontrado"
	@echo ""
	@echo "🚀 Pods:"
	kubectl get pods -l app.kubernetes.io/instance=$(APP_NAME) -n $(NAMESPACE) 2>/dev/null || echo "Nenhum pod encontrado"
	@echo ""
	@echo "🌐 Services:"
	kubectl get svc -l app.kubernetes.io/instance=$(APP_NAME) -n $(NAMESPACE) 2>/dev/null || echo "Nenhum service encontrado"
	@echo ""
	@echo "📈 HPA:"
	kubectl get hpa -l app.kubernetes.io/instance=$(APP_NAME) -n $(NAMESPACE) 2>/dev/null || echo "HPA não encontrado"

helm-logs: ## Mostra logs da aplicação
	@echo "📋 Logs da aplicação:"
	kubectl logs -f deployment/$(APP_NAME) -n $(NAMESPACE)

helm-port-forward: ## Cria port-forward para acessar a aplicação
	@echo "🔌 Criando port-forward para localhost:8080..."
	kubectl port-forward svc/$(APP_NAME) 8080:80 -n $(NAMESPACE)

helm-template: ## Mostra templates gerados pelo Helm
	@echo "🔍 Templates Helm gerados:"
	helm template $(APP_NAME) ./helm/fast-food

helm-lint: ## Valida sintaxe do Helm chart
	@echo "🔍 Validando Helm chart..."
	helm lint ./helm/fast-food

# Aliases para compatibilidade
k8s-deploy: helm-deploy ## Alias para helm-deploy
k8s-delete: helm-delete ## Alias para helm-delete  
k8s-status: helm-status ## Alias para helm-status
k8s-logs: helm-logs ## Alias para helm-logs
k8s-port-forward: helm-port-forward ## Alias para helm-port-forward

test: ## Executa testes
	@echo "🧪 Executando testes..."
	go test ./... -v

swagger-gen: ## Gera documentação Swagger
	@echo "📚 Gerando documentação Swagger..."
	swag init -g cmd/server/main.go -o docs/

clean: ## Limpa arquivos de build
	@echo "🧹 Limpando arquivos..."
	rm -rf bin/
	docker system prune -f

dev: ## Executa em modo desenvolvimento
	@echo "🔧 Iniciando em modo desenvolvimento..."
	go run cmd/server/main.go

# Comandos de ambiente local
env-setup: ## Configura ambiente local
	@echo "⚙️ Configurando ambiente..."
	cp .env.example .env 2>/dev/null || echo "Arquivo .env.example não encontrado"
	@echo "🔑 Edite o arquivo .env com suas configurações"

# Comandos de produção
prod-deploy: helm-deploy ## Deploy para produção
	@echo "🏭 Deploy de produção concluído"

prod-status: helm-status ## Status da produção
	@echo "🏭 Status da produção"

# Comandos utilitários
get-service-url: ## Obtém URL de acesso ao serviço
	@echo "🌐 URLs de acesso:"
	@kubectl get svc $(APP_NAME) -n $(NAMESPACE) -o jsonpath='{.status.loadBalancer.ingress[0].ip}' 2>/dev/null && echo "" || echo "LoadBalancer não disponível"
	@echo "💡 Use 'make helm-port-forward' para acesso local"

helm-upgrade: ## Upgrade do release Helm
	@echo "⬆️ Fazendo upgrade do release..."
	helm upgrade $(APP_NAME) ./helm/fast-food -n $(NAMESPACE)

helm-rollback: ## Rollback do release Helm
	@echo "⬅️ Fazendo rollback do release..."
	helm rollback $(APP_NAME) -n $(NAMESPACE)