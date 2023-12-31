package dev.sunbirdrc.registry.sink;

import dev.sunbirdrc.registry.middleware.util.Constants;
import dev.sunbirdrc.registry.model.DBConnectionInfo;
import dev.sunbirdrc.registry.model.DBConnectionInfoMgr;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

@Component("dbProviderFactory")
public class DBProviderFactory {
	public Map<String, DatabaseProvider> dbProviderInstances = new HashMap<>();

	// This must not be needed. All vars must be sought from DBConnectionInfoMgr only.
	@Autowired
	Environment environment;

	@Autowired
	DBConnectionInfoMgr dbConnectionInfoMgr;

	public DatabaseProvider getInstance(DBConnectionInfo connectionInfo) {
		DatabaseProvider provider = null;
		String dbProvider = environment.getProperty(Constants.DATABASE_PROVIDER);
		String uuidPropertyName = dbConnectionInfoMgr.getUuidPropertyName();
		if (dbProvider == null) {
			throw new RuntimeException("No Database Provider is configured. Please configure a Database Provider");
		}
		if (connectionInfo != null && dbProviderInstances.containsKey(connectionInfo.getShardId())) {
			provider = dbProviderInstances.get(connectionInfo.getShardId());
		} else {
			if (dbProvider.equalsIgnoreCase(Constants.GraphDatabaseProvider.ORIENTDB.getName())) {
				provider = new OrientDBGraphProvider(environment);
				provider.initializeGlobalGraphConfiguration();
			} else if (dbProvider.equalsIgnoreCase(Constants.GraphDatabaseProvider.NEO4J.getName())) {
				provider = new Neo4jGraphProvider(connectionInfo, uuidPropertyName);
			} else if (dbProvider.equalsIgnoreCase(Constants.GraphDatabaseProvider.SQLG.getName())) {
				provider = new SqlgProvider(connectionInfo, uuidPropertyName);
				provider.initializeGlobalGraphConfiguration();
			} else if (dbProvider.equalsIgnoreCase(Constants.GraphDatabaseProvider.TINKERGRAPH.getName())) {
				provider = new TinkerGraphProvider(environment);
				provider.initializeGlobalGraphConfiguration();
			} else if (dbProvider.equalsIgnoreCase(Constants.GraphDatabaseProvider.CASSANDRA.getName())) {
				provider = new JanusGraphStorage(environment, connectionInfo, uuidPropertyName);
				provider.initializeGlobalGraphConfiguration();
			} else {
				throw new RuntimeException("No Database Provider is configured. Please configure a Database Provider");
			}

			if (connectionInfo != null) {
				dbProviderInstances.putIfAbsent(connectionInfo.getShardId(), provider);
			}
		}
		provider.setUuidPropertyName(uuidPropertyName);
		return provider;
	}

}
