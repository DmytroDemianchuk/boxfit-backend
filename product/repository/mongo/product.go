package mongo

import (
	"context"
	"time"

	"github.com/dmytrodemianchuk/boxfit-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"userId"`
	Name        string             `bson:"name"`
	Category    string             `bson:"category"`
	Subcategory string             `bson:"subcategory"`
	Mark        string             `bson:"mark"`
	Variant     string             `bson:"variant"`
	Color       string             `bson:"color"`
	Number      uint16             `bson:"number"`
	Price       uint32             `bson:"price"`
	Image       []byte             `bson:"image"`
	CreatedAt   time.Time          `bson:"create_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type ProductRepository struct {
	db *mongo.Collection
}

func NewProductRepository(db *mongo.Database, collection string) *ProductRepository {
	return &ProductRepository{
		db: db.Collection(collection),
	}
}

func (r ProductRepository) CreateProduct(ctx context.Context, user *models.User, bm *models.Product) error {
	bm.UserID = user.ID
	model := toModel(bm)
	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	bm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r ProductRepository) GetProducts(ctx context.Context, user *models.User) ([]*models.Product, error) {
	uid, _ := primitive.ObjectIDFromHex(user.ID)
	cur, err := r.db.Find(ctx, bson.M{
		"userId": uid,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}
	out := make([]*Product, 0)
	for cur.Next(ctx) {
		user := new(Product)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return toProducts(out), nil
}

func (r ProductRepository) DeleteProduct(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	return err
}

func toModel(p *models.Product) *Product {
	uid, _ := primitive.ObjectIDFromHex(p.UserID)

	return &Product{
		UserID:      uid,
		Name:        p.Name,
		Category:    p.Category,
		Subcategory: p.Subcategory,
		Mark:        p.Mark,
		Variant:     p.Variant,
		Color:       p.Color,
		Number:      p.Number,
		Price:       p.Price,
		Image:       p.Image,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func toProduct(p *Product) *models.Product {
	return &models.Product{
		ID:          p.ID.Hex(),
		UserID:      p.UserID.Hex(),
		Name:        p.Name,
		Category:    p.Category,
		Subcategory: p.Subcategory,
		Mark:        p.Mark,
		Variant:     p.Variant,
		Color:       p.Color,
		Number:      p.Number,
		Price:       p.Price,
		Image:       p.Image,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func toProducts(bs []*Product) []*models.Product {
	out := make([]*models.Product, len(bs))

	for i, p := range bs {
		out[i] = toProduct(p)
	}
	return out
}
