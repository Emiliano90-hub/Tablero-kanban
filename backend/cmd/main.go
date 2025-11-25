package main

import (
	"log"

	http "backend/internal/adapters/in/http/routes"
	adapters "backend/internal/adapters/out/database"
	"backend/internal/application/usecases"
	"backend/internal/application/usecases/board_usecases"
	"backend/internal/application/usecases/column_usecases"
	"backend/internal/application/usecases/task_usecases"
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Cargar configuración
	cfg := config.GetConfig()

	// Conexión a la base de datos
	dbInstance, err := database.NewPostgreConnection(&cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Repositorios
	boardRepo := adapters.NewBoardRepository(dbInstance.Conn)
	columnRepo := adapters.NewColumnRepository(dbInstance.Conn)
	taskRepo := adapters.NewTaskRepository(dbInstance.Conn)

	// Casos de uso
	boardUC := &usecases.BoardUC{
		CreateBoardUC: *board_usecases.NewCreateBoard(boardRepo),
		FindAllBoardsUC: *board_usecases.NewFindAllBoards(boardRepo),
		FindBoardByIdUC: *board_usecases.NewFindBoardByID(boardRepo),
	}
	columnUC := &usecases.ColumnUC{
		CreateColumnUC: *column_usecases.NewCreateColumn(columnRepo),
		FindAllColumnsUC: *column_usecases.NewFindAllColumns(columnRepo),
	}
	taskUC := &usecases.TaskUC{
		CreateTaskUC: *task_usecases.NewCreateTask(taskRepo),
		FindAllTasksUC: *task_usecases.NewFindAllTask(taskRepo),
	}

	// Router Gin
	r := gin.Default()

	// Registrar rutas de Board
	http.SetupRoutes(r, boardUC, columnUC, taskUC)

	// Correr el servidor
	log.Println("🚀 Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
	// // ---------------- Config ----------------
	// cfg := config.GetConfig()

	// // ---------------- DB ----------------
	// db, err := database.NewPostgreConnection(&cfg)
	// if err != nil {
	// 	log.Fatalf("Error connecting to DB: %v", err)
	// }

	// // ---------------- Repos ----------------
	// taskRepo := adapters.NewTaskRepository(db.Conn)
	// columnRepo := adapters.NewColumnRepository(db.Conn)
	// boardRepo := adapters.NewBoardRepository(db.Conn)

	// // ---------------- Services ----------------
	// taskService := services.NewTaskService(taskRepo)
	// columnService := services.NewColumnService(columnRepo)
	// boardService := services.NewBoardService(boardRepo)

	// // ---------------- Router ----------------
	// r := http.SetupRouter(taskService, columnService, boardService)

	// // ---------------- Run ----------------
	// log.Println("🚀 Server running on :8080")
	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatalf("Error running server: %v", err)
	// }

	// // Cargar variables del .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("No se pudo cargar .env, usando variables de entorno del sistema")
	// }

	// // Obtener variables
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")

	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatalf("Error abriendo DB: %v", err)
	// }
	// defer db.Close()

	// // Test de conexión
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("Error conectando a DB: %v", err)
	// }

	// fmt.Println("✅ Conexión exitosa a PostgreSQL!")
}
