package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Document struct {
	ctx    context.Context
	client *mongo.Client
}

func NewDocument(ctx context.Context, client *mongo.Client) (*Document, error) {
	return &Document{
		ctx:    ctx,
		client: client,
	}, nil
}

func ConnectToMongoDB(uri string, ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(10) // Set maximum pool size to 10
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (d *Document) GetContext() context.Context {
	return d.ctx
}

func (d *Document) InsertOne(database string, collection string, document interface{}, validate bool) (*mongo.InsertOneResult, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	opts := options.InsertOne().SetBypassDocumentValidation(validate)

	inserted, err := d.client.Database(database).Collection(collection).InsertOne(
		d.ctx,
		document,
		opts,
	)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return inserted, nil
}

func (d *Document) InsertMany(database string, collection string, documents []interface{}, ordered bool, validate bool) (*mongo.InsertManyResult, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	options := options.InsertMany().SetOrdered(ordered).SetBypassDocumentValidation(validate)
	inserted, err := d.client.Database(database).Collection(collection).InsertMany(d.ctx, documents, options)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)
	return inserted, nil
}

func (d *Document) FindOne(database string, collection string, filter interface{}, fields interface{}) (*mongo.SingleResult, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	options := options.FindOne()

	if fields != nil {
		options.SetProjection(fields)
	}

	singleResult := d.client.Database(database).Collection(collection).FindOne(d.ctx, filter, options)

	session.EndSession(d.ctx)

	return singleResult, singleResult.Err()
}

func (d *Document) Find(
	database string,
	collection string,
	filter interface{},
	fields interface{},
	sort interface{},
	limit int64,
	skip int64,
	usedisk bool,
) (*mongo.Cursor, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	options := options.Find().SetAllowDiskUse(usedisk)

	if fields != nil {
		options.SetProjection(fields)
	}

	if sort != nil {
		options.SetSort(sort)
	}

	if skip > 0 {
		options.SetSkip(skip)
	}

	if limit > 0 {
		options.SetLimit(limit)
	}

	cursor, err := d.client.Database(database).Collection(collection).Find(d.ctx, filter, options)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return cursor, nil
}

func (d *Document) Distinct(database string, collection string, field string, filter interface{}) ([]interface{}, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	result, err := d.client.Database(database).Collection(collection).Distinct(d.ctx, field, filter)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return result, nil
}

func (d *Document) UpdateOne(database string, collection string, filter interface{}, update interface{}, upsert bool, validate bool) (*mongo.UpdateResult, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	options := options.Update().SetUpsert(upsert).SetBypassDocumentValidation(validate)

	updated, err := d.client.Database(database).Collection(collection).UpdateOne(
		d.ctx,
		filter,
		bson.M{
			"$set": update,
		},
		options,
	)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return updated, nil
}

func (d *Document) UpdateMany(database string, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	options := options.Update().SetUpsert(false).SetBypassDocumentValidation(false)

	updated, err := d.client.Database(database).Collection(collection).UpdateMany(
		d.ctx,
		filter,
		bson.M{
			"$set": update,
		},
		options,
	)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return updated, nil
}

func (d *Document) DeleteOne(database string, collection string, filter interface{}) (*mongo.DeleteResult, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	result, err := d.client.Database(database).Collection(collection).DeleteOne(d.ctx, filter)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return result, nil
}

func (d *Document) DeleteMany(database string, collection string, filter interface{}) (*mongo.DeleteResult, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	result, err := d.client.Database(database).Collection(collection).DeleteMany(d.ctx, filter)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return result, nil
}

func (d *Document) Count(database string, collection string, filter interface{}) (int64, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return 0, err
	}

	count, err := d.client.Database(database).Collection(collection).CountDocuments(d.ctx, filter)

	if err != nil {
		return 0, nil
	}

	session.EndSession(d.ctx)

	return count, nil
}

func (d *Document) Aggregate(database string, collection string, pipeline interface{}, usedisk bool, validate bool) (*mongo.Cursor, error) {

	session, err := d.client.StartSession()

	if err != nil {
		return nil, err
	}

	options := options.Aggregate().SetAllowDiskUse(usedisk).SetBypassDocumentValidation(validate)
	result, err := d.client.Database(database).Collection(collection).Aggregate(d.ctx, pipeline, options)

	if err != nil {
		return nil, err
	}

	session.EndSession(d.ctx)

	return result, err
}
